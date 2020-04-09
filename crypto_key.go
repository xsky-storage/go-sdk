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

// CryptoKey is a crypto key
type CryptoKey struct {

	Create time.Time `json:"create,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Size int64 `json:"size,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
