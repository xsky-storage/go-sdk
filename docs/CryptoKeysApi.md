# \CryptoKeysApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCryptoKey**](CryptoKeysApi.md#CreateCryptoKey) | **Post** /crypto-keys/ | 
[**DownloadCryptoKey**](CryptoKeysApi.md#DownloadCryptoKey) | **Get** /crypto-keys/{key_id}/key | 
[**GetCryptoKey**](CryptoKeysApi.md#GetCryptoKey) | **Get** /crypto-keys/{key_id} | 
[**ListCryptoKeys**](CryptoKeysApi.md#ListCryptoKeys) | **Get** /crypto-keys/ | 
[**UpdateCryptoKey**](CryptoKeysApi.md#UpdateCryptoKey) | **Patch** /crypto-keys/{key_id} | 


# **CreateCryptoKey**
> CryptoKeyResp CreateCryptoKey(ctx, name, optional)


Create a crypto key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| crypto key name | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| crypto key name | 
 **key** | **string**| crypto key | 

### Return type

[**CryptoKeyResp**](CryptoKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadCryptoKey**
> string DownloadCryptoKey(ctx, keyId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **keyId** | **int64**| crypto key id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyId** | **int64**| crypto key id | 
 **password** | **string**| password | 

### Return type

**string**

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCryptoKey**
> CryptoKeyResp GetCryptoKey(ctx, keyId)


Get a crypto key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **keyId** | **int64**| crypto key id | 

### Return type

[**CryptoKeyResp**](CryptoKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCryptoKeys**
> CryptoKeysResp ListCryptoKeys(ctx, optional)


List crypto keys

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

[**CryptoKeysResp**](CryptoKeysResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCryptoKey**
> CryptoKeyResp UpdateCryptoKey(ctx, keyId, body)


Update a crypto key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **keyId** | **int64**| crypto key id | 
  **body** | [**CryptoKeyUpdateReq**](CryptoKeyUpdateReq.md)| crypto key info | 

### Return type

[**CryptoKeyResp**](CryptoKeyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

