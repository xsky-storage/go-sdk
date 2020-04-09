# \IdentityPlatformsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityPlatform**](IdentityPlatformsApi.md#CreateIdentityPlatform) | **Post** /identity-platforms/ | 
[**DeleteIdentityPlatform**](IdentityPlatformsApi.md#DeleteIdentityPlatform) | **Delete** /identity-platforms/{identity_platform_id} | 
[**GetIdentityPlatform**](IdentityPlatformsApi.md#GetIdentityPlatform) | **Get** /identity-platforms/{identity_platform_id} | 
[**ListIdentityPlatforms**](IdentityPlatformsApi.md#ListIdentityPlatforms) | **Get** /identity-platforms/ | 
[**RegenerateIdentityKey**](IdentityPlatformsApi.md#RegenerateIdentityKey) | **Post** /identity-platforms/{identity_platform_id}:regenerate | 
[**UpdateIdentityPlatform**](IdentityPlatformsApi.md#UpdateIdentityPlatform) | **Patch** /identity-platforms/{identity_platform_id} | 


# **CreateIdentityPlatform**
> IdentityPlatformResp CreateIdentityPlatform(ctx, body)


Create identity platform

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdentityPlatformCreateReq**](IdentityPlatformCreateReq.md)| identity platform info | 

### Return type

[**IdentityPlatformResp**](IdentityPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIdentityPlatform**
> DeleteIdentityPlatform(ctx, identityPlatformId)


Delete identity platform

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **identityPlatformId** | **int64**| identity platform id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdentityPlatform**
> IdentityPlatformResp GetIdentityPlatform(ctx, identityPlatformId)


Get a identity platform

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **identityPlatformId** | **int64**| identity platform id | 

### Return type

[**IdentityPlatformResp**](IdentityPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIdentityPlatforms**
> IdentityPlatformsResp ListIdentityPlatforms(ctx, optional)


List identity platforms

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

[**IdentityPlatformsResp**](IdentityPlatformsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateIdentityKey**
> IdentityPlatformResp RegenerateIdentityKey(ctx, identityPlatformId)


regenereate a identity platform key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **identityPlatformId** | **int64**| identity platform id | 

### Return type

[**IdentityPlatformResp**](IdentityPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIdentityPlatform**
> IdentityPlatformResp UpdateIdentityPlatform(ctx, identityPlatformId, body)


Update a identity platform

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **identityPlatformId** | **int64**| identity platform id | 
  **body** | [**IdentityPlatformUpdateReq**](IdentityPlatformUpdateReq.md)| identity platform info | 

### Return type

[**IdentityPlatformResp**](IdentityPlatformResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

