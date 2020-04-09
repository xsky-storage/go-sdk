# \AccessTokensApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](AccessTokensApi.md#CreateAccessToken) | **Post** /access-tokens/ | 
[**DeleteAccessToken**](AccessTokensApi.md#DeleteAccessToken) | **Delete** /access-tokens/{access_token_id} | 
[**GetAccessToken**](AccessTokensApi.md#GetAccessToken) | **Get** /access-tokens/{access_token_id} | 
[**ListAccessTokens**](AccessTokensApi.md#ListAccessTokens) | **Get** /access-tokens/ | 
[**RegenerateAccessToken**](AccessTokensApi.md#RegenerateAccessToken) | **Post** /access-tokens/{access_token_id}:regenerate | 
[**UpdateAccessToken**](AccessTokensApi.md#UpdateAccessToken) | **Patch** /access-tokens/{access_token_id} | 
[**ValidateAccessToken**](AccessTokensApi.md#ValidateAccessToken) | **Post** /access-tokens:validate | 


# **CreateAccessToken**
> AccessTokenCreateResp CreateAccessToken(ctx, body)


Create an access token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**AccessTokenCreateReq**](AccessTokenCreateReq.md)| access token info | 

### Return type

[**AccessTokenCreateResp**](AccessTokenCreateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccessToken**
> DeleteAccessToken(ctx, accessTokenId)


delete an access token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accessTokenId** | **int64**| access token id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessToken**
> AccessTokenResp GetAccessToken(ctx, accessTokenId)


get an access Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accessTokenId** | **int64**| access token id | 

### Return type

[**AccessTokenResp**](AccessTokenResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccessTokens**
> AccessTokensResp ListAccessTokens(ctx, optional)


List access tokens

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
 **userId** | **int64**| The id of user access tokens belong to | 
 **role** | **string**| filter access tokens by role | 
 **all** | **bool**| show all access tokens | 

### Return type

[**AccessTokensResp**](AccessTokensResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateAccessToken**
> AccessTokenCreateResp RegenerateAccessToken(ctx, accessTokenId)


regenereate an access token UUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accessTokenId** | **int64**| access token id | 

### Return type

[**AccessTokenCreateResp**](AccessTokenCreateResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccessToken**
> AccessTokenResp UpdateAccessToken(ctx, accessTokenId, body)


update an access token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **accessTokenId** | **int64**| access token id | 
  **body** | [**AccessTokenUpdateReq**](AccessTokenUpdateReq.md)| access token info | 

### Return type

[**AccessTokenResp**](AccessTokenResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateAccessToken**
> AccessTokenResp ValidateAccessToken(ctx, subjectAccessToken)


validate an access token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **subjectAccessToken** | **string**| access token to be validated | 

### Return type

[**AccessTokenResp**](AccessTokenResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

