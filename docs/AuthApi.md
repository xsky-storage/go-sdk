# \AuthApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRSAKey**](AuthApi.md#CreateRSAKey) | **Post** /auth/rsa-keys | 
[**CreateToken**](AuthApi.md#CreateToken) | **Post** /auth/tokens | 
[**GetAuthSecurityPolicy**](AuthApi.md#GetAuthSecurityPolicy) | **Get** /auth/security-policy | 
[**Login**](AuthApi.md#Login) | **Post** /auth/tokens:login | 
[**Logout**](AuthApi.md#Logout) | **Post** /auth/tokens:logout | 
[**UpdateAuthSecurityPolicy**](AuthApi.md#UpdateAuthSecurityPolicy) | **Patch** /auth/security-policy | 


# **CreateRSAKey**
> AuthRsaKeyResp CreateRSAKey(ctx, )


Create RSA key

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthRsaKeyResp**](AuthRSAKeyResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateToken**
> TokenResp CreateToken(ctx, body)


creates temporary token for credentials

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**TokenCreateReq**](TokenCreateReq.md)| the identity credential | 

### Return type

[**TokenResp**](TokenResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthSecurityPolicy**
> AuthSecurityPolicyResp GetAuthSecurityPolicy(ctx, )


setup auth security policy

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuthSecurityPolicyResp**](AuthSecurityPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Login**
> TokenResp Login(ctx, body)


logs in

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**AuthLoginReq**](AuthLoginReq.md)| the identity credential | 

### Return type

[**TokenResp**](TokenResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Logout**
> Logout(ctx, )


logs out

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuthSecurityPolicy**
> AuthSecurityPolicyResp UpdateAuthSecurityPolicy(ctx, body)


update auth security policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**AuthSecurityPolicyUpdateReq**](AuthSecurityPolicyUpdateReq.md)| auth security policy info | 

### Return type

[**AuthSecurityPolicyResp**](AuthSecurityPolicyResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

