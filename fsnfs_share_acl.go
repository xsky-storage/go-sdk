/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

import (
	"time"
)

// FSNFSShareACL defines model of fs nfs share ACL
type FsnfsShareAcl struct {

	AllSquash bool `json:"all_squash,omitempty"`

	Create time.Time `json:"create,omitempty"`

	FsClient *FsClientNestview `json:"fs_client,omitempty"`

	FsClientGroup *FsClientGroupNestview `json:"fs_client_group,omitempty"`

	FsNfsShare *FsnfsShareNestview `json:"fs_nfs_share,omitempty"`

	Id int64 `json:"id,omitempty"`

	Permission string `json:"permission,omitempty"`

	RootSquash bool `json:"root_squash,omitempty"`

	Sync bool `json:"sync,omitempty"`

	Type_ string `json:"type,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
