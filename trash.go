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

// Trash defines the trash attributes.
type Trash struct {

	Create time.Time `json:"create,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Id int64 `json:"id,omitempty"`

	ResourceType string `json:"resource_type,omitempty"`

	Retention string `json:"retention,omitempty"`

	Status string `json:"status,omitempty"`

	TrashResourceNum int64 `json:"trash_resource_num,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
