/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FsQuotaTreeUpdateReqQuotaTree struct {

	// name of quota tree
	Name string `json:"name,omitempty"`

	// size of quota tree
	Size int64 `json:"size,omitempty"`

	// soft quota size of quota tree
	SoftQuotaSize int64 `json:"soft_quota_size,omitempty"`
}