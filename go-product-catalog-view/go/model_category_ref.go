/*
 * Product Catalog View
 *
 * Product Catalog View
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CategoryRef struct {

	// Unique reference of the category.
	BaseType string `json:"@baseType,omitempty"`

	// the class type of the referred category.
	ReferredType string `json:"@referredType,omitempty"`

	// hyperlink reference to the schema describing this resource.
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// Unique reference of the category.
	Type_ string `json:"@type,omitempty"`

	// Unique reference of the category.
	Href string `json:"href,omitempty"`

	// Unique reference of the category.
	Id string `json:"id"`

	// Name of the category.
	Name string `json:"name"`

	// Category version.
	Version string `json:"version,omitempty"`
}
