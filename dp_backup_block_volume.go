/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

// DpBackupBlockVolume is a block volume stored on a site
type DpBackupBlockVolume struct {

	Format int64 `json:"format,omitempty"`

	Name string `json:"name,omitempty"`

	Size int64 `json:"size,omitempty"`

	VolumeName string `json:"volume_name,omitempty"`
}
