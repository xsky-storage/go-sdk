/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type VolumeUpdateReqVolume struct {

	Action string `json:"action,omitempty"`

	BlockSnapshotId int64 `json:"block_snapshot_id,omitempty"`

	Description string `json:"description,omitempty"`

	Flattened bool `json:"flattened,omitempty"`

	Name string `json:"name,omitempty"`

	PerformancePriority int64 `json:"performance_priority,omitempty"`

	Qos *VolumeQosSpec `json:"qos,omitempty"`

	QosEnabled bool `json:"qos_enabled,omitempty"`

	Size int64 `json:"size,omitempty"`
}
