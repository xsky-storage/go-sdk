/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FsGatewayReq struct {

	// host id
	HostId int64 `json:"host_id"`

	// network address id
	NetworkAddressId int64 `json:"network_address_id,omitempty"`
}
