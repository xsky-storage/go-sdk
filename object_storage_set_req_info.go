/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type ObjectStorageSetReqInfo struct {

	// lifecycle end execute time
	LifecycleEndOn string `json:"lifecycle_end_on,omitempty"`

	// lifecycle start execute time
	LifecycleStartOn string `json:"lifecycle_start_on,omitempty"`

	// SecondMergenceEnabled enable second mergence or not
	SecondMergenceEnabled bool `json:"second_mergence_enabled,omitempty"`

	// SecondMergenceEndOn indicates second mergence end time
	SecondMergenceEndOn string `json:"second_mergence_end_on,omitempty"`

	// SecondMergenceStartOn indicates second mergence start time
	SecondMergenceStartOn string `json:"second_mergence_start_on,omitempty"`
}
