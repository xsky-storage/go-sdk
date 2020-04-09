/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type ClusterSetAccessInfoReqCluster struct {

	// access token
	AccessTokenId int64 `json:"access_token_id"`

	// access url
	AccessUrl string `json:"access_url"`
}
