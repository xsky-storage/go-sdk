/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// LifecycleTransition is the details of transition of LifecycleRule
type LifecycleTransition struct {

	StorageClass string `json:"storage_class"`

	Type_ string `json:"type"`
}