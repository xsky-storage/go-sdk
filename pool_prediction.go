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

// PoolPrediction contains pool prediction data
type PoolPrediction struct {

	Create time.Time `json:"create,omitempty"`

	DataKbyte1day int64 `json:"data_kbyte_1day,omitempty"`

	DataKbyte30days int64 `json:"data_kbyte_30days,omitempty"`

	DataKbyte7days int64 `json:"data_kbyte_7days,omitempty"`

	DataKbytePoints []PredictionPoint `json:"data_kbyte_points,omitempty"`

	Id int64 `json:"id,omitempty"`

	UsedKbyte1day int64 `json:"used_kbyte_1day,omitempty"`

	UsedKbyte30days int64 `json:"used_kbyte_30days,omitempty"`

	UsedKbyte7days int64 `json:"used_kbyte_7days,omitempty"`

	UsedKbytePoints []PredictionPoint `json:"used_kbyte_points,omitempty"`
}
