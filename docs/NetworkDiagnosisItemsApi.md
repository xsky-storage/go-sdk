# \NetworkDiagnosisItemsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNetworkDiagnosisItems**](NetworkDiagnosisItemsApi.md#ListNetworkDiagnosisItems) | **Get** /network-diagnosis-items/ | 


# **ListNetworkDiagnosisItems**
> NetworkDiagnosisItemsResp ListNetworkDiagnosisItems(ctx, optional)


List network diagnosis items

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
 **networks** | **string**| network type | 
 **finished** | **bool**| diagnosis item finished or not | 

### Return type

[**NetworkDiagnosisItemsResp**](NetworkDiagnosisItemsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

