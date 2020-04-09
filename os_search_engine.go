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

// OSSearchEngine defines model of OS search engine TODO(penghui): Load balancer group
type OsSearchEngine struct {

	ActionStatus string `json:"action_status,omitempty"`

	AllocatedSize int64 `json:"allocated_size,omitempty"`

	Create time.Time `json:"create,omitempty"`

	GatewayDataSize int64 `json:"gateway_data_size,omitempty"`

	GatewayFlavor *VmFlavorNestview `json:"gateway_flavor,omitempty"`

	GatewayHttpPort int64 `json:"gateway_http_port,omitempty"`

	GatewayTransportPort int64 `json:"gateway_transport_port,omitempty"`

	Id int64 `json:"id,omitempty"`

	SearchS3LoadBalancerGroup *S3LoadBalancerGroupNestview `json:"search_s3_load_balancer_group,omitempty"`

	Size int64 `json:"size,omitempty"`

	Status string `json:"status,omitempty"`

	SyncS3LoadBalancerGroup *S3LoadBalancerGroupNestview `json:"sync_s3_load_balancer_group,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
