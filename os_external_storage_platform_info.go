/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// OSExternalStoragePlatformInfo used for update os external storage class
type OsExternalStoragePlatformInfo struct {

	AccessKey string `json:"access_key,omitempty"`

	Buckets []OsExternalStorageBucketInfo `json:"buckets,omitempty"`

	Endpoint string `json:"endpoint,omitempty"`

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	SecretKey string `json:"secret_key,omitempty"`

	Weight int64 `json:"weight,omitempty"`
}
