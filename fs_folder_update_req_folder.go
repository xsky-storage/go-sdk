/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FsFolderUpdateReqFolder struct {

	// description of folder
	Description string `json:"description,omitempty"`

	// name of folder
	Name string `json:"name,omitempty"`

	// qos of volume
	Qos *VolumeQosSpec `json:"qos,omitempty"`

	// enable or disable the qos
	QosEnabled bool `json:"qos_enabled,omitempty"`

	// size of folder
	Size int64 `json:"size,omitempty"`
}
