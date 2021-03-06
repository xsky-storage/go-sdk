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

// PdRecord defines the response content of protection domain related API
type PdRecord struct {

	// time of creating protection domain
	Create time.Time `json:"create,omitempty"`

	// description of protection domain
	Description string `json:"description,omitempty"`

	// id of protection domain
	Id int64 `json:"id,omitempty"`

	// name of protection domain
	Name string `json:"name,omitempty"`

	// placement node of protection domain
	PlacementNode *PlacementNodeNestview `json:"placement_node,omitempty"`

	// status protection domain
	Status string `json:"status,omitempty"`

	// time of updating protection domain
	Update time.Time `json:"update,omitempty"`
}
