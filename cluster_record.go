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

// ClusterRecord defines the response content of cluster related API
type ClusterRecord struct {

	AccessToken *AccessToken `json:"access_token,omitempty"`

	AccessUrl string `json:"access_url,omitempty"`

	Create time.Time `json:"create,omitempty"`

	DiskLightingMode string `json:"disk_lighting_mode,omitempty"`

	DownOutInterval int64 `json:"down_out_interval,omitempty"`

	ElasticsearchEnabled bool `json:"elasticsearch_enabled,omitempty"`

	FsId string `json:"fs_id,omitempty"`

	Id int64 `json:"id,omitempty"`

	Maintained bool `json:"maintained,omitempty"`

	Name string `json:"name,omitempty"`

	OsGatewayOplogSwitch bool `json:"os_gateway_oplog_switch,omitempty"`

	SnmpEnabled bool `json:"snmp_enabled,omitempty"`

	StatsReservedDays int64 `json:"stats_reserved_days,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`

	Version string `json:"version,omitempty"`

	Samples []ClusterStat `json:"samples,omitempty"`
}
