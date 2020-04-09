# \OsCustomLabelsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSCustomLabel**](OsCustomLabelsApi.md#GetOSCustomLabel) | **Get** /os-custom-labels/{os_custom_label_id} | 
[**ListOSCustomLabels**](OsCustomLabelsApi.md#ListOSCustomLabels) | **Get** /os-custom-labels/ | 


# **GetOSCustomLabel**
> OsCustomLabelResp GetOSCustomLabel(ctx, osCustomLabelId)


Get object storage custom label

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **osCustomLabelId** | **int64**| os custom label id | 

### Return type

[**OsCustomLabelResp**](OSCustomLabelResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSCustomLabels**
> OsCustomLabelsResp ListOSCustomLabels(ctx, optional)


List object storage custom labels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketId** | **int64**| bucket id | 
 **category** | **string**| label category: meta, default, tagging | 
 **limit** | **int64**| paging param | 
 **offset** | **int64**| paging param | 

### Return type

[**OsCustomLabelsResp**](OSCustomLabelsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

