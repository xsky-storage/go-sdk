/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type OsSearchEngineUpdateReqSearchEngine struct {

	GatewayDataSize int64 `json:"gateway_data_size,omitempty"`

	GatewayFlavorId int64 `json:"gateway_flavor_id,omitempty"`
}
