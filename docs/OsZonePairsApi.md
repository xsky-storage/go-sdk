# \OsZonePairsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOSZonePairs**](OsZonePairsApi.md#ListOSZonePairs) | **Get** /os-zone-pairs/ | 


# **ListOSZonePairs**
> OsZonePairsResp ListOSZonePairs(ctx, optional)


list os zone pairs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minClockDiffAbs** | **string**| min clock diff absolute value for zone pairs | 

### Return type

[**OsZonePairsResp**](OSZonePairsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

