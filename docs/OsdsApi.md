# \OsdsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOsd**](OsdsApi.md#CreateOsd) | **Post** /osds/ | 
[**DeleteOsd**](OsdsApi.md#DeleteOsd) | **Delete** /osds/{osd_id} | 
[**GetChunks**](OsdsApi.md#GetChunks) | **Get** /osds/{osd_id}/chunks | 
[**GetOsd**](OsdsApi.md#GetOsd) | **Get** /osds/{osd_id} | 
[**GetOsdPredictions**](OsdsApi.md#GetOsdPredictions) | **Get** /osds/{osd_id}/predictions | 
[**GetOsdSamples**](OsdsApi.md#GetOsdSamples) | **Get** /osds/{osd_id}/samples | 
[**ListOsds**](OsdsApi.md#ListOsds) | **Get** /osds/ | 
[**MaintainOsd**](OsdsApi.md#MaintainOsd) | **Post** /osds/{osd_id}:maintain | 
[**RebuildOsd**](OsdsApi.md#RebuildOsd) | **Post** /osds/{osd_id}:rebuild | 
[**SwitchOsdRole**](OsdsApi.md#SwitchOsdRole) | **Post** /osds/{osd_id}:switch-role | 
[**UnmaintainOsd**](OsdsApi.md#UnmaintainOsd) | **Post** /osds/{osd_id}:unmaintain | 
[**UpdateOsd**](OsdsApi.md#UpdateOsd) | **Patch** /osds/{osd_id} | 


# **CreateOsd**
> OsdResp CreateOsd(ctx, body)


Create osd service on specific disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**OsdCreateReq**](OsdCreateReq.md)| osd info | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOsd**
> OsdResp DeleteOsd(ctx, osdId)


remove an osd from cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osdId** | **int64**| osd id | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChunks**
> ChunksResp GetChunks(ctx, osdId)


get chunks of and osd

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osdId** | **int64**| The id of the osd | 

### Return type

[**ChunksResp**](ChunksResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOsd**
> OsdResp GetOsd(ctx, osdId)


get an osd

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osdId** | **int64**| osd id | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOsdPredictions**
> OsdPredictionsResp GetOsdPredictions(ctx, osdId)


get a osd's prediction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osdId** | **int64**| osd id | 

### Return type

[**OsdPredictionsResp**](OsdPredictionsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOsdSamples**
> OsdSamplesResp GetOsdSamples(ctx, osdId, optional)


get a osd's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osdId** | **int64**| osd id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **osdId** | **int64**| osd id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**OsdSamplesResp**](OsdSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOsds**
> OsdsResp ListOsds(ctx, optional)


List all osds in the cluster

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
 **hostId** | **int64**| host id | 
 **poolId** | **int64**| pool id | 
 **osdGroupId** | **int64**| osd group id | 
 **type_** | **string**| osd type: HDD, SSD, Hybrid | 
 **role** | **string**| osd role: index or data | 
 **withCompound** | **bool**| with compound osd | 
 **withHybrid** | **bool**| with hybrid osd | 
 **cacheDiskId** | **int64**| cache disk id | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**OsdsResp**](OsdsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MaintainOsd**
> OsdResp MaintainOsd(ctx, osdId)


Put osd in maintained status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osdId** | **int64**| osd id | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebuildOsd**
> OsdResp RebuildOsd(ctx, osdId, body)


rebuild an osd from cluster pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osdId** | **int64**| osd id | 
  **body** | [**OsdRebuildReq**](OsdRebuildReq.md)| osd info | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchOsdRole**
> OsdResp SwitchOsdRole(ctx, osdId)


Switch osd role to compound

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osdId** | **int64**| osd id | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnmaintainOsd**
> OsdResp UnmaintainOsd(ctx, osdId)


Put osd out of maintained status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osdId** | **int64**| osd id | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOsd**
> OsdResp UpdateOsd(ctx, osdId, osd)


update osd info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osdId** | **int64**| osd id | 
  **osd** | [**OsdUpdateReq**](OsdUpdateReq.md)| osd info | 

### Return type

[**OsdResp**](OsdResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

