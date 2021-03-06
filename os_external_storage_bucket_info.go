/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// OSExternalStorageBucketInfo contains bucket info of external storage platform
type OsExternalStorageBucketInfo struct {

	Name string `json:"name"`

	Weight int64 `json:"weight"`
}
