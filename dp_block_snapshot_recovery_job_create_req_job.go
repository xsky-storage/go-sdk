/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type DpBlockSnapshotRecoveryJobCreateReqJob struct {

	BackupBlockSnapshot *DpBlockSnapshotRecoveryJobCreateReqJobBackupSnapshot `json:"backup_block_snapshot"`

	BackupBlockVolume *DpBlockSnapshotRecoveryJobCreateReqJobBackupVolume `json:"backup_block_volume"`

	BackupCluster *DpBlockSnapshotRecoveryJobCreateReqJobBackupCluster `json:"backup_cluster"`

	BlockVolume *DpBlockSnapshotRecoveryJobCreateReqJobVolume `json:"block_volume"`

	DpGatewayId int64 `json:"dp_gateway_id"`

	DpSiteId int64 `json:"dp_site_id"`

	ResourceType string `json:"resource_type"`
}