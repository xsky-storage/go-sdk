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

// Partition defines cache partition of disk
type Partition struct {

	ActionStatus string `json:"action_status,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Disk *DiskNestview `json:"disk,omitempty"`

	Id int64 `json:"id,omitempty"`

	OmapByte int64 `json:"omap_byte,omitempty"`

	OmapDevicePath string `json:"omap_device_path,omitempty"`

	Path string `json:"path,omitempty"`

	Size int64 `json:"size,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`

	Used bool `json:"used,omitempty"`

	Uuid string `json:"uuid,omitempty"`

	Version int64 `json:"version,omitempty"`
}
