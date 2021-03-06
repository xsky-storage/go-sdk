/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FsftpShareCreateReqShare struct {

	// folder id
	FsFolderId int64 `json:"fs_folder_id"`

	// access control array
	FsFtpShareAcls []FsftpShareAclReq `json:"fs_ftp_share_acls,omitempty"`

	// gateway group id
	FsGatewayGroupId int64 `json:"fs_gateway_group_id"`

	// quota tree id
	FsQuotaTreeId int64 `json:"fs_quota_tree_id,omitempty"`

	// name of share
	Name string `json:"name,omitempty"`
}
