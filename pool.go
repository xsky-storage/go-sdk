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

// A Pool is a set of temporary objects that may be individually saved and retrieved.  Any item stored in the Pool may be removed automatically at any time without notification. If the Pool holds the only reference when this happens, the item might be deallocated.  A Pool is safe for use by multiple goroutines simultaneously.  Pool's purpose is to cache allocated but unused items for later reuse, relieving pressure on the garbage collector. That is, it makes it easy to build efficient, thread-safe free lists. However, it is not suitable for all free lists.  An appropriate use of a Pool is to manage a group of temporary items silently shared among and potentially reused by concurrent independent clients of a package. Pool provides a way to amortize allocation overhead across many clients.  An example of good use of a Pool is in the fmt package, which maintains a dynamically-sized store of temporary output buffers. The store scales under load (when many goroutines are actively printing) and shrinks when quiescent.  On the other hand, a free list maintained as part of a short-lived object is not a suitable use for a Pool, since the overhead does not amortize well in that scenario. It is more efficient to have such objects implement their own free list.  A Pool must not be copied after first use.
type Pool struct {

	ActionStatus string `json:"action_status,omitempty"`

	BlockVolumeNum int64 `json:"block_volume_num,omitempty"`

	CodingChunkNum int64 `json:"coding_chunk_num,omitempty"`

	Create time.Time `json:"create,omitempty"`

	DataChunkNum int64 `json:"data_chunk_num,omitempty"`

	DefaultManagedVolumeFormat int64 `json:"default_managed_volume_format,omitempty"`

	DeviceType string `json:"device_type,omitempty"`

	DeviceTypeCheckDisabled bool `json:"device_type_check_disabled,omitempty"`

	FailureDomainType string `json:"failure_domain_type,omitempty"`

	Hidden bool `json:"hidden,omitempty"`

	Id int64 `json:"id,omitempty"`

	IoBypassEnabled bool `json:"io_bypass_enabled,omitempty"`

	IoBypassThreshold int64 `json:"io_bypass_threshold,omitempty"`

	Name string `json:"name,omitempty"`

	OsdGroup *OsdGroupNestview `json:"osd_group,omitempty"`

	OsdNum int64 `json:"osd_num,omitempty"`

	PoolId int32 `json:"pool_id,omitempty"`

	PoolMode string `json:"pool_mode,omitempty"`

	PoolName string `json:"pool_name,omitempty"`

	PoolRole string `json:"pool_role,omitempty"`

	PoolType string `json:"pool_type,omitempty"`

	// placement node with the primary replica
	PrimaryPlacementNode *PlacementNodeNestview `json:"primary_placement_node,omitempty"`

	Property *interface{} `json:"property,omitempty"`

	ProtectionDomain *ProtectionDomainNestview `json:"protection_domain,omitempty"`

	Qos *OsdQos `json:"qos,omitempty"`

	ReplicateSize int64 `json:"replicate_size,omitempty"`

	Status string `json:"status,omitempty"`

	Stretched bool `json:"stretched,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
