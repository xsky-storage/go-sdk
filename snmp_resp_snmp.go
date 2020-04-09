/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type SnmpRespSnmp struct {

	// enable snmp or not
	Enabled bool `json:"enabled"`

	// snmp gateways
	Gateways []SnmpGateway `json:"gateways"`

	// snmp receivers
	Receivers []TrapReceiver `json:"receivers"`
}
