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

// OSCustomLabel the struct of object storage custom label
type OsCustomLabel struct {

	Bucket *ObjectStorageBucketNestview `json:"bucket,omitempty"`

	Category string `json:"category,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Status string `json:"status,omitempty"`

	Type_ string `json:"type,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
