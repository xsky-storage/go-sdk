# \OsZonesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectStorageZone**](OsZonesApi.md#CreateObjectStorageZone) | **Post** /os-zones/ | 
[**DeleteObjectStorageZone**](OsZonesApi.md#DeleteObjectStorageZone) | **Delete** /os-zones/{zone_uuid} | 
[**GetObjectStorageZone**](OsZonesApi.md#GetObjectStorageZone) | **Get** /os-zones/{zone_uuid} | 
[**GetObjectStorageZoneSamples**](OsZonesApi.md#GetObjectStorageZoneSamples) | **Get** /os-zones/{zone_uuid}/samples | 
[**ListObjectStorageZones**](OsZonesApi.md#ListObjectStorageZones) | **Get** /os-zones/ | 
[**UpdateOSZonesClockDiff**](OsZonesApi.md#UpdateOSZonesClockDiff) | **Post** /os-zones/{zone_uuid} | 


# **CreateObjectStorageZone**
> ObjectStorageZoneResp CreateObjectStorageZone(ctx, body)


Create a object storage zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ObjectStorageZoneCreateReq**](ObjectStorageZoneCreateReq.md)| zone info | 

### Return type

[**ObjectStorageZoneResp**](ObjectStorageZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteObjectStorageZone**
> ObjectStorageZoneResp DeleteObjectStorageZone(ctx, zoneUuid, optional)


Delete a object storage zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zoneUuid** | **string**| os zone uuid | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneUuid** | **string**| os zone uuid | 
 **force** | **bool**| delete os zone forcefully or not | 

### Return type

[**ObjectStorageZoneResp**](ObjectStorageZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectStorageZone**
> ObjectStorageZoneRecordResp GetObjectStorageZone(ctx, zoneUuid)


Get object storage zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zoneUuid** | **string**| object storage zone uuid | 

### Return type

[**ObjectStorageZoneRecordResp**](ObjectStorageZoneRecordResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectStorageZoneSamples**
> ObjectStorageZoneSamplesResp GetObjectStorageZoneSamples(ctx, zoneUuid, optional)


get an object storage zone's Samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zoneUuid** | **string**| object storage zone uuid | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneUuid** | **string**| object storage zone uuid | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**ObjectStorageZoneSamplesResp**](ObjectStorageZoneSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListObjectStorageZones**
> ObjectStorageZonesRecordResp ListObjectStorageZones(ctx, optional)


List object storage zones

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64**| paging param | 
 **offset** | **int64**| paging param | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 
 **local** | **bool**| local or non-local zones | 
 **master** | **bool**| master or non-master zones | 
 **name** | **string**| name of zone | 

### Return type

[**ObjectStorageZonesRecordResp**](ObjectStorageZonesRecordResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOSZonesClockDiff**
> ObjectStorageZoneResp UpdateOSZonesClockDiff(ctx, zoneUuid, body)


update os zone pairs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zoneUuid** | **string**| os zone uuid | 
  **body** | [**OsZonePairsUpdateReq**](OsZonePairsUpdateReq.md)| zone pairs info | 

### Return type

[**ObjectStorageZoneResp**](ObjectStorageZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

