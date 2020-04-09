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

// SnmpGateway defines the snmp gateway
type SnmpGateway struct {

	AuthPassword string `json:"auth_password,omitempty"`

	AuthProtocol string `json:"auth_protocol,omitempty"`

	ContextEngineId string `json:"context_engine_id,omitempty"`

	ContextName string `json:"context_name,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Host *HostNestview `json:"host,omitempty"`

	Id int64 `json:"id,omitempty"`

	Ip string `json:"ip,omitempty"`

	Port int64 `json:"port,omitempty"`

	PrivPassword string `json:"priv_password,omitempty"`

	PrivProtocol string `json:"priv_protocol,omitempty"`

	// V2c specific
	ReadCommunity string `json:"read_community,omitempty"`

	SecurityEngineId string `json:"security_engine_id,omitempty"`

	SecurityLevel string `json:"security_level,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`

	// V3 specific
	UserName string `json:"user_name,omitempty"`

	Version string `json:"version,omitempty"`

	WriteCommunity string `json:"write_community,omitempty"`
}
