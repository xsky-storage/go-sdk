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

type ClientLunMappingRecord struct {

	BlockVolume *VolumeNestview `json:"block_volume,omitempty"`

	Client *ClientNestview `json:"client,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Host *HostNestview `json:"host,omitempty"`

	Id int64 `json:"id,omitempty"`

	Lun *LunNestview `json:"lun,omitempty"`

	Status string `json:"status,omitempty"`

	Target *TargetNestview `json:"target,omitempty"`

	Update time.Time `json:"update,omitempty"`

	Samples []VolumeStat `json:"samples,omitempty"`

	SessionStatus string `json:"session_status,omitempty"`
}
