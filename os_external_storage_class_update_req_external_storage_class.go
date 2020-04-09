/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type OsExternalStorageClassUpdateReqExternalStorageClass struct {

	Description string `json:"description,omitempty"`

	ExternalStoragePlatforms []OsExternalStoragePlatformInfo `json:"external_storage_platforms,omitempty"`

	Name string `json:"name,omitempty"`

	StoragePattern *SliceStringField `json:"storage_pattern,omitempty"`
}