# S3LoadBalancer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Group** | [***S3LoadBalancerGroupNestview**](S3LoadBalancerGroup_Nestview.md) |  | [optional] [default to null]
**Host** | [***HostNestview**](Host_Nestview.md) |  | [optional] [default to null]
**HttpsPort** | **int64** |  | [optional] [default to null]
**Id** | **int64** |  | [optional] [default to null]
**InterfaceName** | **string** | the interface where vip is bound, exclusive to ip | [optional] [default to null]
**Ip** | **string** | the interface of ip where vip is bound, exclusive to interface_name | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Port** | **int64** |  | [optional] [default to null]
**Roles** | **[]string** |  | [optional] [default to null]
**SearchHttpsPort** | **int64** |  | [optional] [default to null]
**SearchPort** | **int64** |  | [optional] [default to null]
**SslCertificate** | [***SslCertificateNestview**](SSLCertificate_Nestview.md) |  | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**SyncPort** | **int64** |  | [optional] [default to null]
**Update** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Vip** | **string** | designated(master) vip | [optional] [default to null]
**VipMask** | **int64** | mask of vip, default: 32 | [optional] [default to null]
**Vips** | **string** | current effective vips | [optional] [default to null]
**WebConsoleEndpoint** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


