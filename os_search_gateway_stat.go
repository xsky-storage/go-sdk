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

// OSSearchGatewayStat defines model os OS search gateway stat
type OsSearchGatewayStat struct {

	CpuUtil float64 `json:"cpu_util,omitempty"`

	Create time.Time `json:"create,omitempty"`

	MemUsagePercent float64 `json:"mem_usage_percent,omitempty"`

	ReadBandwidthKbyte float64 `json:"read_bandwidth_kbyte,omitempty"`

	ReadIops float64 `json:"read_iops,omitempty"`

	ReadLatencyUs float64 `json:"read_latency_us,omitempty"`

	WriteBandwidthKbyte float64 `json:"write_bandwidth_kbyte,omitempty"`

	WriteIops float64 `json:"write_iops,omitempty"`

	WriteLatencyUs float64 `json:"write_latency_us,omitempty"`
}
