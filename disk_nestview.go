/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type DiskNestview struct {

	Device string `json:"device,omitempty"`

	DiskType string `json:"disk_type,omitempty"`

	Id int64 `json:"id,omitempty"`
}
