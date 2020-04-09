# \OsReplicationZonesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOSReplicationZone**](OsReplicationZonesApi.md#CreateOSReplicationZone) | **Post** /os-replication-zones/ | 
[**DeleteOSReplicationZone**](OsReplicationZonesApi.md#DeleteOSReplicationZone) | **Delete** /os-replication-zones/{zone_uuid} | 
[**GetOSReplicationZone**](OsReplicationZonesApi.md#GetOSReplicationZone) | **Get** /os-replication-zones/{zone_uuid} | 
[**GetOSReplicationZoneSamples**](OsReplicationZonesApi.md#GetOSReplicationZoneSamples) | **Get** /os-replication-zones/{zone_uuid}/samples | 
[**ListOSReplicationZones**](OsReplicationZonesApi.md#ListOSReplicationZones) | **Get** /os-replication-zones/ | 


# **CreateOSReplicationZone**
> OsReplicationZoneResp CreateOSReplicationZone(ctx, body)


Create a os replication zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**OsReplicationZoneCreateReq**](OsReplicationZoneCreateReq.md)| os replication zone info | 

### Return type

[**OsReplicationZoneResp**](OSReplicationZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOSReplicationZone**
> OsReplicationZoneResp DeleteOSReplicationZone(ctx, zoneUuid)


Delete a os replication zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zoneUuid** | **string**| os replication zone uuid | 

### Return type

[**OsReplicationZoneResp**](OSReplicationZoneResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOSReplicationZone**
> OsReplicationZoneRecordResp GetOSReplicationZone(ctx, zoneUuid)


Get a os replication zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zoneUuid** | **string**| os replication zone uuid | 

### Return type

[**OsReplicationZoneRecordResp**](OSReplicationZoneRecordResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOSReplicationZoneSamples**
> OsReplicationZoneSamplesResp GetOSReplicationZoneSamples(ctx, zoneUuid, optional)


get an os replication zone's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **zoneUuid** | **string**| os replication zone uuid | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneUuid** | **string**| os replication zone uuid | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**OsReplicationZoneSamplesResp**](OSReplicationZoneSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSReplicationZones**
> OsReplicationZoneRecordsResp ListOSReplicationZones(ctx, optional)


List os replication zones

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
 **marker** | **string**| paging param | 
 **replicationUuid** | **string**| os replication uuid | 
 **osZoneUuid** | **string**| os zone uuid | 

### Return type

[**OsReplicationZoneRecordsResp**](OSReplicationZoneRecordsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

