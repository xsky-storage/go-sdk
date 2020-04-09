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

// MappingGroup defines the mapping group of volumes and client groups
type MappingGroup struct {

	AccessPath *AccessPathNestview `json:"access_path,omitempty"`

	ClientGroup *ClientGroupNestview `json:"client_group,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Id int64 `json:"id,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
