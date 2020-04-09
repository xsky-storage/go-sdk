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

// OsdGroupRecord combine OsdGroup and OsdGroupStat as API response
type OsdGroupRecord struct {

	ActionStatus string `json:"action_status,omitempty"`

	Create time.Time `json:"create,omitempty"`

	DeviceType string `json:"device_type,omitempty"`

	DeviceTypeCheckDisabled bool `json:"device_type_check_disabled,omitempty"`

	FailureDomainType string `json:"failure_domain_type,omitempty"`

	Id int64 `json:"id,omitempty"`

	OsdAsyncRecoveryMaxUpdates int64 `json:"osd_async_recovery_max_updates,omitempty"`

	OsdFullRatio float64 `json:"osd_full_ratio,omitempty"`

	OsdNum int64 `json:"osd_num,omitempty"`

	Pools []PoolNestview `json:"pools,omitempty"`

	Qos *OsdQos `json:"qos,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`

	Samples []OsdGroupStat `json:"samples,omitempty"`
}
