/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type OsBucketAddReplicationPathsReqBucket struct {

	OsReplicationPathUuids []string `json:"os_replication_path_uuids,omitempty"`

	OsReplicationPaths []OsReplicationPathReq `json:"os_replication_paths,omitempty"`
}