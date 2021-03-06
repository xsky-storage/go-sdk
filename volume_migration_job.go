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

// VolumeMigrationJob defines the volume migration job record
type VolumeMigrationJob struct {

	Create time.Time `json:"create,omitempty"`

	DestPool *PoolNestview `json:"dest_pool,omitempty"`

	FinishedAt time.Time `json:"finished_at,omitempty"`

	Host *HostNestview `json:"host,omitempty"`

	Id int64 `json:"id,omitempty"`

	RemainingTime int64 `json:"remaining_time,omitempty"`

	SourcePool *PoolNestview `json:"source_pool,omitempty"`

	StartedAt time.Time `json:"started_at,omitempty"`

	Status string `json:"status,omitempty"`

	Stripes int64 `json:"stripes,omitempty"`

	Task *Task `json:"task,omitempty"`

	Update time.Time `json:"update,omitempty"`

	Volume *VolumeNestview `json:"volume,omitempty"`
}
