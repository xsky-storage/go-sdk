/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type ObjectStorageKeyCreateReqKey struct {

	AccessKey string `json:"access_key"`

	SecretKey string `json:"secret_key,omitempty"`

	UserId int64 `json:"user_id"`
}