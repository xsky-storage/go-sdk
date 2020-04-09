# \OsGatewaysApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGateway**](OsGatewaysApi.md#CreateGateway) | **Post** /os-gateways/ | 
[**DeleteGateway**](OsGatewaysApi.md#DeleteGateway) | **Delete** /os-gateways/{gateway_id} | 
[**GetGateway**](OsGatewaysApi.md#GetGateway) | **Get** /os-gateways/{gateway_id} | 
[**GetGatewaySamples**](OsGatewaysApi.md#GetGatewaySamples) | **Get** /os-gateways/{gateway_id}/samples | 
[**ListGateways**](OsGatewaysApi.md#ListGateways) | **Get** /os-gateways/ | 
[**UpdateGateway**](OsGatewaysApi.md#UpdateGateway) | **Put** /os-gateways/{gateway_id} | 


# **CreateGateway**
> ObjectStorageGatewayResp CreateGateway(ctx, body)


Create s3 gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ObjectStorageGatewayCreateReq**](ObjectStorageGatewayCreateReq.md)| gateway info | 

### Return type

[**ObjectStorageGatewayResp**](ObjectStorageGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGateway**
> ObjectStorageGatewayResp DeleteGateway(ctx, gatewayId)


Delete s3 gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| gateway id | 

### Return type

[**ObjectStorageGatewayResp**](ObjectStorageGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGateway**
> ObjectStorageGatewayResp GetGateway(ctx, gatewayId)


Get s3 gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| gateway id | 

### Return type

[**ObjectStorageGatewayResp**](ObjectStorageGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewaySamples**
> ObjectStorageGatewaySamplesResp GetGatewaySamples(ctx, gatewayId, optional)


Get s3 gateway's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| gateway id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayId** | **int64**| gateway id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**ObjectStorageGatewaySamplesResp**](ObjectStorageGatewaySamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGateways**
> ObjectStorageGatewaysResp ListGateways(ctx, optional)


List s3 gateways

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

### Return type

[**ObjectStorageGatewaysResp**](ObjectStorageGatewaysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGateway**
> ObjectStorageGatewayResp UpdateGateway(ctx, gatewayId, body)


Update s3 gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| gateway id | 
  **body** | [**ObjectStorageGatewayUpdateReq**](ObjectStorageGatewayUpdateReq.md)| gateway info | 

### Return type

[**ObjectStorageGatewayResp**](ObjectStorageGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

