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

// CloudVolume defines volumes of the cloud platform
type CloudVolume struct {

	BlockVolume *VolumeNestview `json:"block_volume,omitempty"`

	CloudDatacenter *CloudDatacenterNestview `json:"cloud_datacenter,omitempty"`

	CloudPlatform *CloudPlatformNestview `json:"cloud_platform,omitempty"`

	CloudVolumeId string `json:"cloud_volume_id,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Id int64 `json:"id,omitempty"`

	MultiAttach bool `json:"multi_attach,omitempty"`

	Name string `json:"name,omitempty"`

	Type_ string `json:"type,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
