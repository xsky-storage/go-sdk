# \ConfsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfItem**](ConfsApi.md#DeleteConfItem) | **Delete** /confs/{type}/{key} | 
[**GetConfItem**](ConfsApi.md#GetConfItem) | **Get** /confs/{type}/{key} | 
[**ListConfItems**](ConfsApi.md#ListConfItems) | **Get** /confs/{type} | 
[**SetConfItem**](ConfsApi.md#SetConfItem) | **Put** /confs/ | 
[**SetConfItems**](ConfsApi.md#SetConfItems) | **Put** /confs/{type} | 


# **DeleteConfItem**
> DeleteConfItem(ctx, type_, key)


delete conf item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **type_** | **string**| filter the type of conf item | 
  **key** | **string**| filter the key of conf item | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfItem**
> ConfItemResp GetConfItem(ctx, type_, key)


get conf item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **type_** | **string**| filter the type of conf | 
  **key** | **string**| filter the key of conf item | 

### Return type

[**ConfItemResp**](ConfItemResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConfItems**
> ConfItemsResp ListConfItems(ctx, type_, optional)


List conf items

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **type_** | **string**| filter the type of conf | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| filter the type of conf | 
 **limit** | **int64**| paging param | 
 **offset** | **int64**| paging param | 

### Return type

[**ConfItemsResp**](ConfItemsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetConfItem**
> ConfItem SetConfItem(ctx, body)


set conf item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ConfItemSetReq**](ConfItemSetReq.md)| conf item | 

### Return type

[**ConfItem**](ConfItem.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetConfItems**
> SetConfItems(ctx, type_, body)


set conf items

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **type_** | **string**| filter the type of conf | 
  **body** | [**ConfItemsSetReq**](ConfItemsSetReq.md)| conf items | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

