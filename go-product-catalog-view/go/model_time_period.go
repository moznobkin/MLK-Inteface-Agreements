/*
 * Product Catalog View
 *
 * Product Catalog View
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type TimePeriod struct {

	// End date of the catalog
	EndDateTime time.Time `json:"endDateTime,omitempty"`

	// Start date of the catalog
	StartDateTime time.Time `json:"startDateTime,omitempty"`
}
