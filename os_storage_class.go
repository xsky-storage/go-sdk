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

// OSStorageClass is the model of os_storage_class
type OsStorageClass struct {

	ActivePoolIds []int64 `json:"active_pool_ids,omitempty"`

	Class int32 `json:"class,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Description string `json:"description,omitempty"`

	Id int64 `json:"id,omitempty"`

	InactivePoolIds []int64 `json:"inactive_pool_ids,omitempty"`

	Name string `json:"name,omitempty"`

	OsPolicy *ObjectStoragePolicy `json:"os_policy,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
