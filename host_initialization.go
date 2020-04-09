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

// HostInitialization is the model of host initialization
type HostInitialization struct {

	AdminIps []string `json:"admin_ips,omitempty"`

	Create time.Time `json:"create,omitempty"`

	DisableFirewalld bool `json:"disable_firewalld,omitempty"`

	HostnamePrefix string `json:"hostname_prefix,omitempty"`

	Id int64 `json:"id,omitempty"`

	Message string `json:"message,omitempty"`

	SetHostname bool `json:"set_hostname,omitempty"`

	SshUser string `json:"ssh_user,omitempty"`

	Status string `json:"status,omitempty"`
}