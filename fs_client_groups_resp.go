/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FsClientGroupsResp struct {

	// file storage client groups
	FsClientGroups []FsClientGroup `json:"fs_client_groups"`
}
