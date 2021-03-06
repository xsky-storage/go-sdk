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

// ObjectStorageBucket is the model of object_storage_bucket
type ObjectStorageBucket struct {

	ActionStatus string `json:"action_status,omitempty"`

	AllUserPermission string `json:"all_user_permission,omitempty"`

	AuthUserPermission string `json:"auth_user_permission,omitempty"`

	BucketPolicy string `json:"bucket_policy,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Flag *interface{} `json:"flag,omitempty"`

	Id int64 `json:"id,omitempty"`

	Lifecycle *ObjectStorageLifecycle `json:"lifecycle,omitempty"`

	LogDeliveryPermission string `json:"log_delivery_permission,omitempty"`

	LoggingEnabled bool `json:"logging_enabled,omitempty"`

	LoggingGrants []OsBucketLoggingGrant `json:"logging_grants,omitempty"`

	LoggingOwner *ObjectStorageUserNestview `json:"logging_owner,omitempty"`

	LoggingPrefix string `json:"logging_prefix,omitempty"`

	LoggingSuspended bool `json:"logging_suspended,omitempty"`

	LoggingTargetBucket *ObjectStorageBucketNestview `json:"logging_target_bucket,omitempty"`

	MetadataSearchEnabled bool `json:"metadata_search_enabled,omitempty"`

	Name string `json:"name,omitempty"`

	NfsClientNum int64 `json:"nfs_client_num,omitempty"`

	OsReplicationPathNum int64 `json:"os_replication_path_num,omitempty"`

	OsReplicationZoneNum int64 `json:"os_replication_zone_num,omitempty"`

	OsZone *ObjectStorageZoneNestview `json:"os_zone,omitempty"`

	OsZoneUuid string `json:"os_zone_uuid,omitempty"`

	Owner *ObjectStorageUserNestview `json:"owner,omitempty"`

	OwnerPermission string `json:"owner_permission,omitempty"`

	Policy *ObjectStoragePolicyNestview `json:"policy,omitempty"`

	PolicyEnabled bool `json:"policy_enabled,omitempty"`

	QuotaMaxObjects int64 `json:"quota_max_objects,omitempty"`

	QuotaMaxSize int64 `json:"quota_max_size,omitempty"`

	RemoteCluster *RemoteClusterNestview `json:"remote_cluster,omitempty"`

	// NOTE: Use name of bucket as replication uuid for simplicity
	ReplicationUuid string `json:"replication_uuid,omitempty"`

	Shards int64 `json:"shards,omitempty"`

	SpecificationObjects int64 `json:"specification_objects,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`

	Virtual bool `json:"virtual,omitempty"`
}
