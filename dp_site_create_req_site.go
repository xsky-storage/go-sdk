/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type DpSiteCreateReqSite struct {

	// site address
	Address string `json:"address,omitempty"`

	// extra config
	Config *DpSiteConfig `json:"config"`

	// crypto key
	CryptoKeyId int64 `json:"crypto_key_id,omitempty"`

	// site name
	Name string `json:"name"`

	// platform of site
	Platform string `json:"platform"`

	// protection type
	ProtectionType string `json:"protection_type"`

	// remote cluster
	RemoteClusterId int64 `json:"remote_cluster_id,omitempty"`

	// service of site
	Service string `json:"service"`
}