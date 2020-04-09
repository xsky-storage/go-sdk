/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FsGatewayGroupRemoveGatewaysReqGatewayGroup struct {

	// file storage gateway ids
	FsGatewayIds []int64 `json:"fs_gateway_ids"`
}
