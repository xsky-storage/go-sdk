# FsGatewayGroupCreateReqGatewayGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | description of gateway group | [optional] [default to null]
**Encoding** | **string** | ftp encoding format, default is utf8 | [optional] [default to null]
**FsGateways** | [**[]FsGatewayReq**](FSGatewayReq.md) | file storage gateways list | [default to null]
**Name** | **string** | name of gateway group | [default to null]
**NfsVersions** | **[]string** | nfs versions of nfs supported | [optional] [default to null]
**Security** | **string** | smb security type | [optional] [default to null]
**Smb1Enabled** | **bool** | smb version 1.0 enabled | [optional] [default to null]
**SmbPorts** | **[]int64** | smb ports | [optional] [default to null]
**Types** | **[]string** | types of supported (smb, nfs, ftp) | [default to null]
**Vip** | **string** | virtual ip of gateway group | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


