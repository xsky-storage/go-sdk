/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FssmbShareAclReq struct {

	// id of user group
	FsUserGroupId int64 `json:"fs_user_group_id,omitempty"`

	// id of user group
	FsUserGroupName string `json:"fs_user_group_name,omitempty"`

	// id of user
	FsUserId int64 `json:"fs_user_id,omitempty"`

	// id of user
	FsUserName string `json:"fs_user_name,omitempty"`

	// id of user group
	Id int64 `json:"id,omitempty"`

	// readonly or readwrite access
	Permission string `json:"permission,omitempty"`

	// type of share acl
	Type_ string `json:"type,omitempty"`
}
