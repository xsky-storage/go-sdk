/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// SnapshotBackupConfig contains config for snapshot backup
type SnapshotBackupConfig struct {

	AccessKey string `json:"access_key,omitempty"`

	AppId string `json:"app_id,omitempty"`

	Bucket string `json:"bucket,omitempty"`

	MountPath string `json:"mount_path,omitempty"`

	Password string `json:"password,omitempty"`

	Port int64 `json:"port,omitempty"`

	Region string `json:"region,omitempty"`

	SecretKey string `json:"secret_key,omitempty"`

	Username string `json:"username,omitempty"`
}
