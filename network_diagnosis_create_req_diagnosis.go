/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type NetworkDiagnosisCreateReqDiagnosis struct {

	// ids of host
	HostIds []int64 `json:"host_ids"`

	// host scope
	HostScope string `json:"host_scope"`

	// network types
	Networks []string `json:"networks"`
}