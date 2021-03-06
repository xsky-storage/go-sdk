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

// NetworkDiagnosis defines model of network diagnosis
type NetworkDiagnosis struct {

	Create time.Time `json:"create,omitempty"`

	DiagnoseActiveNum int64 `json:"diagnose_active_num,omitempty"`

	DiagnoseErrorNum int64 `json:"diagnose_error_num,omitempty"`

	DiagnoseNum int64 `json:"diagnose_num,omitempty"`

	DiagnoseWarningNum int64 `json:"diagnose_warning_num,omitempty"`

	DiagnosingNum int64 `json:"diagnosing_num,omitempty"`

	HostNum int64 `json:"host_num,omitempty"`

	Id int64 `json:"id,omitempty"`

	Networks []string `json:"networks,omitempty"`

	NotDiagnoseNum int64 `json:"not_diagnose_num,omitempty"`

	Progress float64 `json:"progress,omitempty"`

	Status string `json:"status,omitempty"`

	Update time.Time `json:"update,omitempty"`
}
