/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FsftpShareAclReq struct {

	// enable creating files
	CreateEnabled bool `json:"create_enabled,omitempty"`

	// enable deleting files
	DeleteEnabled bool `json:"delete_enabled,omitempty"`

	// max bandwidth of downloading
	DownloadBandwidth int64 `json:"download_bandwidth,omitempty"`

	// enable downloading files
	DownloadEnabled bool `json:"download_enabled,omitempty"`

	// id of user group
	FsUserGroupId int64 `json:"fs_user_group_id,omitempty"`

	// id of user
	FsUserId int64 `json:"fs_user_id,omitempty"`

	// id of user group
	Id int64 `json:"id,omitempty"`

	// enable listing files
	ListEnabled bool `json:"list_enabled,omitempty"`

	// enable renaming files
	RenameEnabled bool `json:"rename_enabled,omitempty"`

	// type of share acl
	Type_ string `json:"type,omitempty"`

	// max bandwidth of uploading
	UploadBandwidth int64 `json:"upload_bandwidth,omitempty"`

	// enable uploading files
	UploadEnabled bool `json:"upload_enabled,omitempty"`
}