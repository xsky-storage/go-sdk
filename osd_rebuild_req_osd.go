/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type OsdRebuildReqOsd struct {

	// new data disk id
	NewDiskId int64 `json:"new_disk_id,omitempty"`

	// new size of omap partition
	NewOmapByte int64 `json:"new_omap_byte,omitempty"`

	// new cache partition id
	NewPartitionId int64 `json:"new_partition_id,omitempty"`

	// new read cache size in bytes
	NewReadCacheSize int64 `json:"new_read_cache_size,omitempty"`
}
