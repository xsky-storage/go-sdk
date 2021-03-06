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

// NetworkInterfaceRecord combine NetworkInterface and NetworkInterfaceStat as API response
type NetworkInterfaceRecord struct {

	BondingMode string `json:"bonding_mode,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Host *HostNestview `json:"host,omitempty"`

	Id int64 `json:"id,omitempty"`

	LinkDetected bool `json:"link_detected,omitempty"`

	MacAddress string `json:"mac_address,omitempty"`

	MasterNetworkInterface *NetworkInterfaceNestview `json:"master_network_interface,omitempty"`

	Megabits int64 `json:"megabits,omitempty"`

	Name string `json:"name,omitempty"`

	Operstate string `json:"operstate,omitempty"`

	// ethernet or bond
	Type_ string `json:"type,omitempty"`

	Update time.Time `json:"update,omitempty"`

	Samples []NetworkInterfaceStat `json:"samples,omitempty"`
}
