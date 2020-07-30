package swagger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// IRefLoader is an abstraction for loading references for composite
type IRefLoader interface {
	LoadProductOfferingPrice(id string) (*ProductOfferingPrice, []error)
	LoadProductOffering(id string) (*ProductOffering, []error)
}

// FileRefLoader is a file based implementation of IRefLoader
type FileRefLoader struct {
	baseDir string
}

// LoadProductOfferingPrice loads full ProductOfferingPrice
func (l FileRefLoader) LoadProductOfferingPrice(id string) (*ProductOfferingPrice, []error) {
	filename := fmt.Sprintf("%s/ProductOfferingPrice/ProductOfferingPrice-%s.json", l.baseDir, id)
	fs, err := os.Open(filename)
	if err != nil {
		return nil, []error{err}
	}

	var p ProductOfferingPrice

	err = json.NewDecoder(fs).Decode(&p)
	if err != nil {
		return nil, []error{err}
	}
	return &p, nil
}

func (l FileRefLoader) load(id string, typename string) (*os.File, error) {

	filename := fmt.Sprintf("%s/%s/%s-%s.json", l.baseDir, typename, typename, id)
	fs, err := os.Open(filename)
	if err != nil {
		log.Println("File does not exist ", filename)
	}
	return fs, err

}

// LoadProductSpecification loads ProductSpecification
func (l FileRefLoader) LoadProductSpecification(id string) (*ProductSpecification, []error) {
	var p *ProductSpecification
	fs, err := l.load(id, "ProductSpecification")
	if err != nil {
		return nil, []error{err}
	}

	err = json.NewDecoder(fs).Decode(&p)
	if err != nil {
		return nil, []error{err}
	}
	errors := make([]error, 0)
	for i, pSpec := range p.ProductSpecCharacteristic {
		full, errs := l.LoadCharacteristics(pSpec.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
		if full != nil {
			p.ProductSpecCharacteristic[i] = *full
		}
	}
	return p, errors

}

// LoadProductSpecification loads ProductSpecification
func (l FileRefLoader) LoadCharacteristics(id string) (*SpecCharacteristic, []error) {

	var p *SpecCharacteristic
	fs, err := l.load(id, "Characteristic")
	if err != nil {
		return nil, []error{err}
	}

	err = json.NewDecoder(fs).Decode(&p)
	if err != nil {
		return nil, []error{err}
	}
	return p, nil

}

// LoadProductOffering loads full ProductOffering
func (l FileRefLoader) LoadProductOffering(id string) (*ProductOffering, []error) {
	var errs []error
	filename := fmt.Sprintf("%s/ProductOffering/ProductOffering-%s.json", l.baseDir, id)
	fmt.Printf(filename)
	fs, err := os.Open(filename)
	if err != nil {
		return nil, []error{err}
	}

	var p ProductOffering

	err = json.NewDecoder(fs).Decode(&p)
	if err != nil {
		return nil, []error{err}
	}
	errors := make([]error, 0)
	for i, pop := range p.ProductOfferingPrice {
		full, errs := l.LoadProductOfferingPrice(pop.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
		if full != nil {
			p.ProductOfferingPrice[i] = *full
		}

	}
	if p.ProductSpecification != nil {
		p.ProductSpecification, errs = l.LoadProductSpecification(p.ProductSpecification.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	for i, bundle := range p.BundledProductOffering {
		full, errs := l.LoadProductOffering(bundle.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
		if full != nil {
			p.BundledProductOffering[i] = *full
		}
	}

	return &p, errors
}

// LoadDir loads the directory
func LoadDir(baseDir string) ([]*ProductOffering, []error) {

	errors := make([]error, 0)
	productOfferings := make([]*ProductOffering, 0)
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if f.IsDir() {
			p, e := LoadDir(baseDir + "/" + f.Name())
			productOfferings = append(productOfferings, p...)
			errors = append(errors, e...)
		} else {
			if strings.HasPrefix(f.Name(), "ProductOffering-") {
				var l IRefLoader = FileRefLoader{baseDir: filepath.Dir(baseDir)}
				id := strings.Replace(strings.Replace(f.Name(), "ProductOffering-", "", 1), ".json", "", 1)
				po, errs := l.LoadProductOffering(id)
				if errs != nil {
					errors = append(errors, errs...)
				}
				if po != nil {
					productOfferings = append(productOfferings, po)
				}
			}
		}
	}
	return productOfferings, errors
}
