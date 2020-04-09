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

// OSReplicationZone defines models of os replication zone.
type OsReplicationZone struct {

	Create time.Time `json:"create,omitempty"`

	Dirty bool `json:"dirty,omitempty"`

	OsBucket *ObjectStorageBucketNestview `json:"os_bucket,omitempty"`

	OsBucketOwnerCluster *RemoteClusterNestview `json:"os_bucket_owner_cluster,omitempty"`

	OsBucketOwnerZone *ObjectStorageZoneNestview `json:"os_bucket_owner_zone,omitempty"`

	OsRemotePolicy *OsRemotePolicyNestview `json:"os_remote_policy,omitempty"`

	OsRemotePolicyUuid string `json:"os_remote_policy_uuid,omitempty"`

	OsReplicationPathNum int64 `json:"os_replication_path_num,omitempty"`

	OsZone *ObjectStorageZoneNestview `json:"os_zone,omitempty"`

	OsZoneUuid string `json:"os_zone_uuid,omitempty"`

	ReplicationUuid string `json:"replication_uuid,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}
