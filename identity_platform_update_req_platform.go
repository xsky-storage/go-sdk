/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type IdentityPlatformUpdateReqPlatform struct {

	Enabled bool `json:"enabled,omitempty"`

	// name of identity platform
	Name string `json:"name,omitempty"`

	Url string `json:"url,omitempty"`
}
