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

// NetworkInterfaceStat define the statistics of network interface
type NetworkInterfaceStat struct {

	Create time.Time `json:"create,omitempty"`

	DropPps float64 `json:"drop_pps,omitempty"`

	ErrorPps float64 `json:"error_pps,omitempty"`

	RxBandwidthKbyte float64 `json:"rx_bandwidth_kbyte,omitempty"`

	RxDropPps float64 `json:"rx_drop_pps,omitempty"`

	RxErrorPps float64 `json:"rx_error_pps,omitempty"`

	RxPps float64 `json:"rx_pps,omitempty"`

	TxBandwidthKbyte float64 `json:"tx_bandwidth_kbyte,omitempty"`

	TxDropPps float64 `json:"tx_drop_pps,omitempty"`

	TxErrorPps float64 `json:"tx_error_pps,omitempty"`

	TxPps float64 `json:"tx_pps,omitempty"`
}