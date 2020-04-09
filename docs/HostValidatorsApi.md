# \HostValidatorsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostValidator**](HostValidatorsApi.md#CreateHostValidator) | **Post** /host-validators/ | 
[**GetHostValidator**](HostValidatorsApi.md#GetHostValidator) | **Get** /host-validators/{host_validator_id} | 


# **CreateHostValidator**
> HostValidatorResp CreateHostValidator(ctx, body)


Create host validator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**HostValidatorCreateReq**](HostValidatorCreateReq.md)| host validator info | 

### Return type

[**HostValidatorResp**](HostValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHostValidator**
> HostValidatorResp GetHostValidator(ctx, hostValidatorId)


Get host validator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hostValidatorId** | **int64**| host validator id | 

### Return type

[**HostValidatorResp**](HostValidatorResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

