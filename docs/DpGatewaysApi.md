# \DpGatewaysApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpGateway**](DpGatewaysApi.md#CreateDpGateway) | **Post** /dp-gateways/ | 
[**DeleteDpGateway**](DpGatewaysApi.md#DeleteDpGateway) | **Delete** /dp-gateways/{gateway_id} | 
[**GetDpGateway**](DpGatewaysApi.md#GetDpGateway) | **Get** /dp-gateways/{gateway_id} | 
[**ListDpGateways**](DpGatewaysApi.md#ListDpGateways) | **Get** /dp-gateways/ | 
[**UpdateDpGateway**](DpGatewaysApi.md#UpdateDpGateway) | **Patch** /dp-gateways/{gateway_id} | 


# **CreateDpGateway**
> DpGatewayResp CreateDpGateway(ctx, body)


Create a dp gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**DpGatewayCreateReq**](DpGatewayCreateReq.md)| dp gateway info | 

### Return type

[**DpGatewayResp**](DpGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDpGateway**
> DpGatewayResp DeleteDpGateway(ctx, gatewayId)


Delete dp gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| dp gateway id | 

### Return type

[**DpGatewayResp**](DpGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDpGateway**
> DpGatewayResp GetDpGateway(ctx, gatewayId)


Get dp gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| dp gateway id | 

### Return type

[**DpGatewayResp**](DpGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDpGateways**
> DpGatewaysResp ListDpGateways(ctx, optional)


List dp gateways

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

[**DpGatewaysResp**](DpGatewaysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDpGateway**
> DpGatewayResp UpdateDpGateway(ctx, gatewayId, body)


Update a dp gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **gatewayId** | **int64**| dp gateway id | 
  **body** | [**DpGatewayUpdateReq**](DpGatewayUpdateReq.md)| dp gateway info | 

### Return type

[**DpGatewayResp**](DpGatewayResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

