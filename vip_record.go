/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

import (
	"time"
)

// VIPRecord represents doc returned by API.
type VipRecord struct {

	ActionStatus string `json:"action_status,omitempty"`

	Create time.Time `json:"create,omitempty"`

	Id int64 `json:"id,omitempty"`

	Ip string `json:"ip,omitempty"`

	MacAddress string `json:"mac_address,omitempty"`

	Mask int64 `json:"mask,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`

	VipGroup *VipGroupNestview `json:"vip_group,omitempty"`

	VirtualRouterId int64 `json:"virtual_router_id,omitempty"`

	CurrentVipInstance *VipInstanceNestview `json:"current_vip_instance,omitempty"`

	// Copy these two fields from corresponding fields of VIP to avoid JSON marshal recursion.
	DefaultVipInstance *VipInstanceNestview `json:"default_vip_instance,omitempty"`
}