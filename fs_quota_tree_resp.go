/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FsQuotaTreeResp struct {

	// file storage quota tree
	FsQuotaTree *FsQuotaTree `json:"fs_quota_tree"`
}
