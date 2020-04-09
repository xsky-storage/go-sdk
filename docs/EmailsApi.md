# \EmailsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailConfig**](EmailsApi.md#GetEmailConfig) | **Get** /emails/config | 
[**SendEmail**](EmailsApi.md#SendEmail) | **Post** /emails/ | 
[**UpdateEmailConfig**](EmailsApi.md#UpdateEmailConfig) | **Patch** /emails/config | 


# **GetEmailConfig**
> EmailConfigResp GetEmailConfig(ctx, )


setup email configs

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EmailConfigResp**](EmailConfigResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendEmail**
> EmailResp SendEmail(ctx, body)


send email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**EmailSendReq**](EmailSendReq.md)| email info | 

### Return type

[**EmailResp**](EmailResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmailConfig**
> EmailConfigResp UpdateEmailConfig(ctx, body)


setup email configs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**EmailConfigUpdateReq**](EmailConfigUpdateReq.md)| email config info | 

### Return type

[**EmailConfigResp**](EmailConfigResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

