/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type OsSearchGatewayReq struct {

	// host id
	HostId int64 `json:"host_id"`

	PoolId int64 `json:"pool_id"`
}