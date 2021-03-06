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

// FSClient defines model of file storage client
type FsClient struct {

	Create time.Time `json:"create,omitempty"`

	FsClientGroupNum int64 `json:"fs_client_group_num,omitempty"`

	FsClientGroups []FsClientGroupNestview `json:"fs_client_groups,omitempty"`

	Id int64 `json:"id,omitempty"`

	Ip string `json:"ip,omitempty"`

	Name string `json:"name,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
