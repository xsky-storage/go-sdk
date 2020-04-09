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

// TrapReceiver defines the snmp trap receiver
type TrapReceiver struct {

	Create time.Time `json:"create,omitempty"`

	Id int64 `json:"id,omitempty"`

	Ip string `json:"ip,omitempty"`

	MsgLevel string `json:"msg_level,omitempty"`

	Port int64 `json:"port,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
