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

// ObjectStorageUserRecord is used to combine ObjectStorageUser and ObjectStorageUserStat
type ObjectStorageUserRecord struct {

	BucketNum int64 `json:"bucket_num,omitempty"`

	BucketQuotaMaxObjects int64 `json:"bucket_quota_max_objects,omitempty"`

	BucketQuotaMaxSize int64 `json:"bucket_quota_max_size,omitempty"`

	Create time.Time `json:"create,omitempty"`

	DisplayName string `json:"display_name,omitempty"`

	Email string `json:"email,omitempty"`

	Id int64 `json:"id,omitempty"`

	LocationConstraintEnabled bool `json:"location_constraint_enabled,omitempty"`

	MaxBuckets int64 `json:"max_buckets,omitempty"`

	Name string `json:"name,omitempty"`

	OpMask string `json:"op_mask,omitempty"`

	Parent *ObjectStorageUserNestview `json:"parent,omitempty"`

	Policy *ObjectStoragePolicyNestview `json:"policy,omitempty"`

	PolicyPollingEnabled bool `json:"policy_polling_enabled,omitempty"`

	Properties *OsUserProperties `json:"properties,omitempty"`

	Status string `json:"status,omitempty"`

	Suspended bool `json:"suspended,omitempty"`

	Update time.Time `json:"update,omitempty"`

	UserQuotaMaxObjects int64 `json:"user_quota_max_objects,omitempty"`

	UserQuotaMaxSize int64 `json:"user_quota_max_size,omitempty"`

	Samples []ObjectStorageUserStat `json:"samples,omitempty"`
}
