/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type PoolRuleReq struct {

	PlacementNodeId int64 `json:"placement_node_id,omitempty"`

	ReplicateNum int64 `json:"replicate_num,omitempty"`
}