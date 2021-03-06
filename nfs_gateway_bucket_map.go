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

// NFSGatewayBucketMap defines nfs gateway bucket map @grpc-models-proto
type NfsGatewayBucketMap struct {

	ActionStatus string `json:"action_status,omitempty"`

	Bucket *ObjectStorageBucketNestview `json:"bucket,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Id int64 `json:"id,omitempty"`

	Key *ObjectStorageKeyNestview `json:"key,omitempty"`

	MountClients string `json:"mount_clients,omitempty"`

	MountNum int64 `json:"mount_num,omitempty"`

	NfsGateway *NfsGatewayNestview `json:"nfs_gateway,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
