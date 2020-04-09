/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type ObjectStorageUserCreateReqInfo struct {

	BucketQuotaMaxObjects int64 `json:"bucket_quota_max_objects,omitempty"`

	BucketQuotaMaxSize int64 `json:"bucket_quota_max_size,omitempty"`

	DisplayName string `json:"display_name,omitempty"`

	Email string `json:"email,omitempty"`

	Keys []ObjectStorageKey `json:"keys,omitempty"`

	LocationConstraintEnabled bool `json:"location_constraint_enabled,omitempty"`

	MaxBuckets int64 `json:"max_buckets,omitempty"`

	Name string `json:"name"`

	OpMask string `json:"op_mask,omitempty"`

	PolicyId int64 `json:"policy_id,omitempty"`

	PolicyPollingEnabled bool `json:"policy_polling_enabled,omitempty"`

	Properties *OsUserPropertiesReq `json:"properties,omitempty"`

	UserQuotaMaxObjects int64 `json:"user_quota_max_objects,omitempty"`

	UserQuotaMaxSize int64 `json:"user_quota_max_size,omitempty"`
}