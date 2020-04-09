# \AlertGroupsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertGroup**](AlertGroupsApi.md#CreateAlertGroup) | **Post** /alert-groups/ | 
[**DeleteAlertGroup**](AlertGroupsApi.md#DeleteAlertGroup) | **Delete** /alert-groups/{group_id} | 
[**GetAlertGroup**](AlertGroupsApi.md#GetAlertGroup) | **Get** /alert-groups/{group_id} | 
[**ListAlertGroups**](AlertGroupsApi.md#ListAlertGroups) | **Get** /alert-groups/ | 
[**UpdateAlertGroup**](AlertGroupsApi.md#UpdateAlertGroup) | **Patch** /alert-groups/{group_id} | 


# **CreateAlertGroup**
> AlertGroupResp CreateAlertGroup(ctx, body)


create alert group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**AlertGroupCreateReq**](AlertGroupCreateReq.md)| alert group | 

### Return type

[**AlertGroupResp**](AlertGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlertGroup**
> DeleteAlertGroup(ctx, groupId)


delete alert group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| the id of alert group | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertGroup**
> AlertGroupResp GetAlertGroup(ctx, groupId)


get alert group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| the id of alert group | 

### Return type

[**AlertGroupResp**](AlertGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAlertGroups**
> AlertGroupsResp ListAlertGroups(ctx, optional)


List all alert groups

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

[**AlertGroupsResp**](AlertGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAlertGroup**
> AlertGroupResp UpdateAlertGroup(ctx, groupId, body)


update alert group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| the id of alert group | 
  **body** | [**AlertGroupUpdateReq**](AlertGroupUpdateReq.md)| alert group | 

### Return type

[**AlertGroupResp**](AlertGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

