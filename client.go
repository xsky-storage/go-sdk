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

type Client struct {

	ClientGroup *ClientGroupNestview `json:"client_group,omitempty"`

	Code string `json:"code,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Id int64 `json:"id,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
