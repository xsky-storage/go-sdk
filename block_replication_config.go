/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// BlockReplicationConfig contains config for block replication
type BlockReplicationConfig struct {

	Destination string `json:"destination,omitempty"`

	LocalGateways []string `json:"local_gateways,omitempty"`

	RemoteGateways []string `json:"remote_gateways,omitempty"`
}
