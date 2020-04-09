/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type PoolCreateReqPool struct {

	CodingChunkNum int64 `json:"coding_chunk_num,omitempty"`

	DataChunkNum int64 `json:"data_chunk_num,omitempty"`

	FailureDomainType string `json:"failure_domain_type,omitempty"`

	Name string `json:"name,omitempty"`

	OsdIds []int64 `json:"osd_ids,omitempty"`

	PoolRole string `json:"pool_role,omitempty"`

	PoolType string `json:"pool_type,omitempty"`

	PrimaryPlacementNodeId int64 `json:"primary_placement_node_id,omitempty"`

	ProtectionDomainId int64 `json:"protection_domain_id,omitempty"`

	Ruleset []PoolRuleReq `json:"ruleset,omitempty"`

	SharedPoolId int64 `json:"shared_pool_id,omitempty"`

	Size int64 `json:"size,omitempty"`

	Stretched bool `json:"stretched,omitempty"`
}
