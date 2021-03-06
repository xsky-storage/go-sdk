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

// VMFlavor defines flavor of vm
type VmFlavor struct {

	Cpu int64 `json:"cpu,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Id int64 `json:"id,omitempty"`

	MemoryKbyte int64 `json:"memory_kbyte,omitempty"`

	Name string `json:"name,omitempty"`

	RootDiskSize int64 `json:"root_disk_size,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
