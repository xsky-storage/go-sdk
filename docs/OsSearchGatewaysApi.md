# \OsSearchGatewaysApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSSearchGatewaySamples**](OsSearchGatewaysApi.md#GetOSSearchGatewaySamples) | **Get** /os-search-gateways/{gateway_id}/samples | 
[**GetOSSearchGateways**](OsSearchGatewaysApi.md#GetOSSearchGateways) | **Get** /os-search-gateways/{gateway_id} | 
[**ListOSSearchGateways**](OsSearchGatewaysApi.md#ListOSSearchGateways) | **Get** /os-search-gateways/ | 


# **GetOSSearchGatewaySamples**
> OsSearchGatewaySamplesResp GetOSSearchGatewaySamples(ctx, gatewayId, optional)


get a os search gateway's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| os search gateway id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64**| os search gateway id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**OsSearchGatewaySamplesResp**](OSSearchGatewaySamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOSSearchGateways**
> OsSearchGatewayResp GetOSSearchGateways(ctx, gatewayId)


Get OS Search Gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| os search gateway id | 

### Return type

[**OsSearchGatewayResp**](OSSearchGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSSearchGateways**
> OsSearchGatewaysResp ListOSSearchGateways(ctx, optional)


List OS Search Gateways

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **osSearchEngineId** | **int64**| os search engine id | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**OsSearchGatewaysResp**](OSSearchGatewaysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

