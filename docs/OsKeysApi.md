# \OsKeysApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKey**](OsKeysApi.md#CreateKey) | **Post** /os-keys/ | 
[**DeleteKey**](OsKeysApi.md#DeleteKey) | **Delete** /os-keys/{key_id} | 
[**GetKey**](OsKeysApi.md#GetKey) | **Get** /os-keys/{key_id} | 
[**ListKeys**](OsKeysApi.md#ListKeys) | **Get** /os-keys/ | 
[**UpdateKey**](OsKeysApi.md#UpdateKey) | **Patch** /os-keys/{key_id} | 


# **CreateKey**
> ObjectStorageKeyResp CreateKey(ctx, body)


Create new object storage key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ObjectStorageKeyCreateReq**](ObjectStorageKeyCreateReq.md)| key info | 

### Return type

[**ObjectStorageKeyResp**](ObjectStorageKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteKey**
> ObjectStorageKeyResp DeleteKey(ctx, keyId)


Delete object storage key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **keyId** | **int64**| key id | 

### Return type

[**ObjectStorageKeyResp**](ObjectStorageKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKey**
> ObjectStorageKeyResp GetKey(ctx, keyId)


get object storage key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **keyId** | **int64**| user id | 

### Return type

[**ObjectStorageKeyResp**](ObjectStorageKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListKeys**
> ObjectStorageKeysResp ListKeys(ctx, optional)


List object storage keys

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
 **userId** | **int64**| object storage user id | 

### Return type

[**ObjectStorageKeysResp**](ObjectStorageKeysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateKey**
> ObjectStorageKeyResp UpdateKey(ctx, keyId, body)


Update object storage key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **keyId** | **int64**| key id | 
  **body** | [**ObjectStorageKeyUpdateReq**](ObjectStorageKeyUpdateReq.md)| key info | 

### Return type

[**ObjectStorageKeyResp**](ObjectStorageKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

