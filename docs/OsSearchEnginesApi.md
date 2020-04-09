# \OsSearchEnginesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOSSearchGateways**](OsSearchEnginesApi.md#AddOSSearchGateways) | **Post** /os-search-engines/{os_search_engine_id}:add-os-search-gateways | 
[**ChangeOSSearchEngine**](OsSearchEnginesApi.md#ChangeOSSearchEngine) | **Patch** /os-search-engines/{os_search_engine_id} | 
[**CreateOSSearchEngine**](OsSearchEnginesApi.md#CreateOSSearchEngine) | **Post** /os-search-engines/ | 
[**DeleteOSSearchEngine**](OsSearchEnginesApi.md#DeleteOSSearchEngine) | **Delete** /os-search-engines/{os_search_engine_id} | 
[**GetOSSearchEngine**](OsSearchEnginesApi.md#GetOSSearchEngine) | **Get** /os-search-engines/{os_search_engine_id} | 
[**GetOSSearchEngineSamples**](OsSearchEnginesApi.md#GetOSSearchEngineSamples) | **Get** /os-search-engines/{os_search_engine_id}/samples | 
[**ListOSSearchEngines**](OsSearchEnginesApi.md#ListOSSearchEngines) | **Get** /os-search-engines/ | 
[**RemoveOSSearchGateways**](OsSearchEnginesApi.md#RemoveOSSearchGateways) | **Post** /os-search-engines/{os_search_engine_id}:remove-os-search-gateways | 
[**StartOSSearchEngine**](OsSearchEnginesApi.md#StartOSSearchEngine) | **Post** /os-search-engines/{os_search_engine_id}:start | 
[**StopOSSearchEngine**](OsSearchEnginesApi.md#StopOSSearchEngine) | **Post** /os-search-engines/{os_search_engine_id}:stop | 


# **AddOSSearchGateways**
> OsSearchEngineResp AddOSSearchGateways(ctx, osSearchEngineId, body)


Create new OS Search gateways on OS search engine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osSearchEngineId** | **int64**| OS search engine id | 
  **body** | [**OsSearchEngineAddGatewaysReq**](OsSearchEngineAddGatewaysReq.md)| os search gateways info | 

### Return type

[**OsSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeOSSearchEngine**
> OsSearchEngineResp ChangeOSSearchEngine(ctx, osSearchEngineId, body)


change OS search engine falvor or data size

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osSearchEngineId** | **int64**| OS search engine id | 
  **body** | [**OsSearchEngineUpdateReq**](OsSearchEngineUpdateReq.md)| os search gateways info | 

### Return type

[**OsSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOSSearchEngine**
> OsSearchEngineResp CreateOSSearchEngine(ctx, body)


Create OS search engine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**OsSearchEngineCreateReq**](OsSearchEngineCreateReq.md)| OS search engine info | 

### Return type

[**OsSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOSSearchEngine**
> OsSearchEngineResp DeleteOSSearchEngine(ctx, osSearchEngineId)


delete OS search engine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osSearchEngineId** | **int64**| OS search engine id | 

### Return type

[**OsSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOSSearchEngine**
> OsSearchEngineResp GetOSSearchEngine(ctx, osSearchEngineId)


Get a OS search engine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osSearchEngineId** | **int64**| OS search engine id | 

### Return type

[**OsSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOSSearchEngineSamples**
> OsSearchEngineSamplesResp GetOSSearchEngineSamples(ctx, osSearchEngineId, optional)


get an object storage search engine's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osSearchEngineId** | **int64**| OS search engine id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **osSearchEngineId** | **int64**| OS search engine id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**OsSearchEngineSamplesResp**](OSSearchEngineSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSSearchEngines**
> OsSearchEnginesResp ListOSSearchEngines(ctx, optional)


List OS search engine

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

### Return type

[**OsSearchEnginesResp**](OSSearchEnginesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOSSearchGateways**
> OsSearchEngineResp RemoveOSSearchGateways(ctx, osSearchEngineId, body)


remove OS search gateways from OS search engine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osSearchEngineId** | **int64**| OS search engine id | 
  **body** | [**OsSearchEngineRemoveGatewaysReq**](OsSearchEngineRemoveGatewaysReq.md)| os search gateways info | 

### Return type

[**OsSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartOSSearchEngine**
> OsSearchEngineResp StartOSSearchEngine(ctx, osSearchEngineId)


start OS search gateways from OS search engine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osSearchEngineId** | **int64**| OS search engine id | 

### Return type

[**OsSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopOSSearchEngine**
> OsSearchEngineResp StopOSSearchEngine(ctx, osSearchEngineId)


stop OS search gateways from OS search engine

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osSearchEngineId** | **int64**| OS search engine id | 

### Return type

[**OsSearchEngineResp**](OSSearchEngineResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

