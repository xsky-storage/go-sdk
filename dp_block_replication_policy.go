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

// DpBlockReplicationPolicy protects a block volume by sync replication
type DpBlockReplicationPolicy struct {

	Create time.Time `json:"create,omitempty"`

	DpSite *DpSiteNestview `json:"dp_site,omitempty"`

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Status string `json:"status,omitempty"`

	TimeoutSeconds int64 `json:"timeout_seconds,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
