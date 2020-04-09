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

// ClusterStat defines the statistics metric of a ceph cluster
type ClusterStat struct {

	ActualKbyte int64 `json:"actual_kbyte,omitempty"`

	Create time.Time `json:"create,omitempty"`

	DataKbyte int64 `json:"data_kbyte,omitempty"`

	DegradedPercent float64 `json:"degraded_percent,omitempty"`

	ErrorKbyte int64 `json:"error_kbyte,omitempty"`

	HealthyPercent float64 `json:"healthy_percent,omitempty"`

	OsDownBandwidthKbyte float64 `json:"os_down_bandwidth_kbyte,omitempty"`

	OsDownIops float64 `json:"os_down_iops,omitempty"`

	OsMergeSpeed int64 `json:"os_merge_speed,omitempty"`

	OsUpBandwidthKbyte float64 `json:"os_up_bandwidth_kbyte,omitempty"`

	OsUpIops float64 `json:"os_up_iops,omitempty"`

	ReadBandwidthKbyte float64 `json:"read_bandwidth_kbyte,omitempty"`

	ReadIops float64 `json:"read_iops,omitempty"`

	ReadLatencyUs float64 `json:"read_latency_us,omitempty"`

	RecoveryBandwidthKbyte float64 `json:"recovery_bandwidth_kbyte,omitempty"`

	RecoveryIops float64 `json:"recovery_iops,omitempty"`

	RecoveryPercent float64 `json:"recovery_percent,omitempty"`

	TotalKbyte int64 `json:"total_kbyte,omitempty"`

	UnavailablePercent float64 `json:"unavailable_percent,omitempty"`

	UsedKbyte int64 `json:"used_kbyte,omitempty"`

	WriteBandwidthKbyte float64 `json:"write_bandwidth_kbyte,omitempty"`

	WriteIops float64 `json:"write_iops,omitempty"`

	WriteLatencyUs float64 `json:"write_latency_us,omitempty"`
}
