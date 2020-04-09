/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type RemoteClusterUpdateReqRemoteCluster struct {

	// access token of remote cluster
	AccessToken string `json:"access_token"`

	// URL of remote cluster
	Url string `json:"url"`
}
