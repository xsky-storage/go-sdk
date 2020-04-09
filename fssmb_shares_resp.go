/*
 * XMS API
 *
 * XMS is the controller of distributed storage system
 *
 * API version: SDS_4.2.000.0.200302
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package xmsclient

type FssmbSharesResp struct {

	// fs smb shares
	FsSmbShares []FssmbShare `json:"fs_smb_shares"`
}