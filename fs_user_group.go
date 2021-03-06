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

// FSUserGroup defines model of file storage user group
type FsUserGroup struct {

	ActionStatus string `json:"action_status,omitempty"`

	Create time.Time `json:"create,omitempty"`

	FsUsers []FsUserNestview `json:"fs_users,omitempty"`

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Update time.Time `json:"update,omitempty"`

	UserNum int64 `json:"user_num,omitempty"`
}
