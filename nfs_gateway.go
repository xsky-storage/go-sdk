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

// NFSGateway defines nfs gateway @grpc-models-proto
type NfsGateway struct {

	BucketNum int64 `json:"bucket_num,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Description string `json:"description,omitempty"`

	GatewayIp string `json:"gateway_ip,omitempty"`

	GatewayName string `json:"gateway_name,omitempty"`

	Host *HostNestview `json:"host,omitempty"`

	Id int64 `json:"id,omitempty"`

	MountClients string `json:"mount_clients,omitempty"`

	MountNum int64 `json:"mount_num,omitempty"`

	Name string `json:"name,omitempty"`

	Port int64 `json:"port,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`
}