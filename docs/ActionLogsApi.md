# \ActionLogsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListActionLogs**](ActionLogsApi.md#ListActionLogs) | **Get** /action-logs/ | 


# **ListActionLogs**
> ActionLogsResp ListActionLogs(ctx, optional)


List action logs

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
 **resourceType** | **string**| resource type of action logs | 
 **status** | **string**| status of action logs | 
 **userId** | **int64**| user id of action logs | 
 **createBegin** | **string**| create begin timestamp | 
 **createEnd** | **string**| create end timestamp | 
 **q** | **string**| query param of search | 
 **relatedResource** | **string**| related resource info of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**ActionLogsResp**](ActionLogsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

