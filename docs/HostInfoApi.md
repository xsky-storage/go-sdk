# \HostInfoApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHostInfo**](HostInfoApi.md#GetHostInfo) | **Get** /host-info/ | 


# **GetHostInfo**
> HostInfoResp GetHostInfo(ctx, ip)


scan host info from specified ip

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ip** | **string**| host ip | 

### Return type

[**HostInfoResp**](HostInfoResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

