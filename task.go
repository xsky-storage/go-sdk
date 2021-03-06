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

// Task defines a task will be run on a host @grpc-models-proto
type Task struct {

	Create time.Time `json:"create,omitempty"`

	Data string `json:"data,omitempty"`

	ErrorMsg string `json:"error_msg,omitempty"`

	Execute time.Time `json:"execute,omitempty"`

	Finish time.Time `json:"finish,omitempty"`

	Host *HostNestview `json:"host,omitempty"`

	HostRole string `json:"host_role,omitempty"`

	Id int64 `json:"id,omitempty"`

	Priority int64 `json:"priority,omitempty"`

	Retried bool `json:"retried,omitempty"`

	Retry int64 `json:"retry,omitempty"`

	RetryType string `json:"retry_type,omitempty"`

	Status string `json:"status,omitempty"`

	Type_ string `json:"type,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
