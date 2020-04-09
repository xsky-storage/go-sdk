/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// OSSample contains os samples
type OsSample struct {

	DeleteObjCategory *OsBucketUsageCategory `json:"delete_obj_category,omitempty"`

	GetObjCategory *OsBucketUsageCategory `json:"get_obj_category,omitempty"`

	ListBucketCategory *OsBucketUsageCategory `json:"list_bucket_category,omitempty"`

	Objects int64 `json:"objects,omitempty"`

	PostObjCategory *OsBucketUsageCategory `json:"post_obj_category,omitempty"`

	PutObjCategory *OsBucketUsageCategory `json:"put_obj_category,omitempty"`

	UsedCapacityBytes map[string]int64 `json:"used_capacity_bytes,omitempty"`
}
