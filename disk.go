/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

import (
	"time"
)

// Disk defines the available disk @grpc-models-proto
type Disk struct {

	ActionStatus string `json:"action_status,omitempty"`

	// size of disk
	Bytes int64 `json:"bytes,omitempty"`

	CacheCreate time.Time `json:"cache_create,omitempty"`

	ChannelId string `json:"channel_id,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Device string `json:"device,omitempty"`

	DiskType string `json:"disk_type,omitempty"`

	DriverType string `json:"driver_type,omitempty"`

	EnclosureId string `json:"enclosure_id,omitempty"`

	Host *HostNestview `json:"host,omitempty"`

	Id int64 `json:"id,omitempty"`

	// used as cache disk
	IsCache bool `json:"is_cache,omitempty"`

	// used as root disk
	IsRoot bool `json:"is_root,omitempty"`

	LightingStatus string `json:"lighting_status,omitempty"`

	Model string `json:"model,omitempty"`

	PartitionNum int64 `json:"partition_num,omitempty"`

	Partitions []Partition `json:"partitions,omitempty"`

	PowerSafe bool `json:"power_safe,omitempty"`

	RotationRate string `json:"rotation_rate,omitempty"`

	Rotational bool `json:"rotational,omitempty"`

	Serial string `json:"serial,omitempty"`

	SlotId string `json:"slot_id,omitempty"`

	SmartAttrs []SmartAttr `json:"smart_attrs,omitempty"`

	SsdLifeLeft int64 `json:"ssd_life_left,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`

	Used bool `json:"used,omitempty"`

	Wwid string `json:"wwid,omitempty"`
}
