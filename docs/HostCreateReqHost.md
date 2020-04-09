# HostCreateReqHost

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminIp** | **string** | admin ip | [default to null]
**Description** | **string** | host description | [optional] [default to null]
**GatewayIps** | **[]string** | gateway ips for s3 | [optional] [default to null]
**MetaDevice** | **string** | meta device for docker | [optional] [default to null]
**PrivateIp** | **string** | cluster private ip for internal data access | [optional] [default to null]
**ProtectionDomainId** | **int64** | deprecated | [optional] [default to null]
**PublicIp** | **string** | public ip for outside data access | [optional] [default to null]
**Roles** | **[]string** | host roles: admin,monitor,block_storage_gateway,file_storage_gateway,s3_gateway,nfs_gateway | [optional] [default to null]
**Type_** | **string** | storage server or storage client | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


