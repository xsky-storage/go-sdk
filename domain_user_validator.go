/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

import (
	"time"
)

// DomainUserValidator defines the validator for domain users
type DomainUserValidator struct {

	Create time.Time `json:"create,omitempty"`

	Id int64 `json:"id,omitempty"`

	Members *interface{} `json:"members,omitempty"`

	Report *interface{} `json:"report,omitempty"`

	Status string `json:"status,omitempty"`

	Type_ string `json:"type,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
