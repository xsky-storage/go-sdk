/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// VolumeDpBlockBackupPolicyMapping represent the mapping relationship of volume and dp block backup policies
type VolumeDpBlockBackupPolicyMapping struct {

	DpBlockBackupPolicy *DpBlockBackupPolicyNestview `json:"dp_block_backup_policy,omitempty"`

	Id int64 `json:"id,omitempty"`

	Status string `json:"status,omitempty"`

	Volume *VolumeNestview `json:"volume,omitempty"`
}