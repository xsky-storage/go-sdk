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

// PredictionPoint is a point in a prediction graph
type PredictionPoint struct {

	Create time.Time `json:"create,omitempty"`

	UsedKbyte int64 `json:"used_kbyte,omitempty"`
}
