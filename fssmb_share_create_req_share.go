/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FssmbShareCreateReqShare struct {

	// default acl status
	AclInherited bool `json:"acl_inherited,omitempty"`

	// case sensitive
	CaseSensitive bool `json:"case_sensitive,omitempty"`

	// folder id
	FsFolderId int64 `json:"fs_folder_id"`

	// gateway group id
	FsGatewayGroupId int64 `json:"fs_gateway_group_id"`

	// quota tree id
	FsQuotaTreeId int64 `json:"fs_quota_tree_id,omitempty"`

	// access control array
	FsSmbShareAcls []FssmbShareAclReq `json:"fs_smb_share_acls,omitempty"`

	// name of share
	Name string `json:"name,omitempty"`

	// recycle status
	Recycled bool `json:"recycled,omitempty"`
}
