/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type OsZoneLockCreateReqLock struct {

	// key of lock, for example name of os bucket
	Key string `json:"key"`

	// resource type of lock, including os_bucket
	ResourceType string `json:"resource_type"`

	// value of lock, for example action of os bucket
	Value string `json:"value"`
}
