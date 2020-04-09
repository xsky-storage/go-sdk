/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type VipGroupCreateReqVipGroup struct {

	Network string `json:"network"`

	Preempt bool `json:"preempt,omitempty"`

	ResourceId int64 `json:"resource_id"`

	ResourceType string `json:"resource_type"`

	Vips []VipGroupCreateReqVipGroupViPsElt `json:"vips"`
}
