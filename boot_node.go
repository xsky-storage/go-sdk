/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// BootNode defines the status of boot node
type BootNode struct {

	AdminNetwork string `json:"admin_network"`

	GatewayNetwork string `json:"gateway_network"`

	Host *HostNestview `json:"host,omitempty"`

	PrivateNetwork string `json:"private_network"`

	PublicNetwork string `json:"public_network"`

	Status string `json:"status,omitempty"`
}
