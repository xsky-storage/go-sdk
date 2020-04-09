# \EmailGroupsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailGroup**](EmailGroupsApi.md#CreateEmailGroup) | **Post** /email-groups/ | 
[**DeleteEmailGroup**](EmailGroupsApi.md#DeleteEmailGroup) | **Delete** /email-groups/{group_id} | 
[**GetEmailGroup**](EmailGroupsApi.md#GetEmailGroup) | **Get** /email-groups/{group_id} | 
[**ListEmailGroups**](EmailGroupsApi.md#ListEmailGroups) | **Get** /email-groups/ | 
[**UpdateEmailGroup**](EmailGroupsApi.md#UpdateEmailGroup) | **Put** /email-groups/{group_id} | 


# **CreateEmailGroup**
> EmailGroupResp CreateEmailGroup(ctx, body)


create email group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**EmailGroupUpdateReq**](EmailGroupUpdateReq.md)| email group | 

### Return type

[**EmailGroupResp**](EmailGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEmailGroup**
> DeleteEmailGroup(ctx, groupId)


delete email group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| the id of email group | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailGroup**
> EmailGroupResp GetEmailGroup(ctx, groupId)


get email group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| the id of email group | 

### Return type

[**EmailGroupResp**](EmailGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEmailGroups**
> EmailGroupsResp ListEmailGroups(ctx, optional)


List all email groups

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
 **alertGroupId** | **int64**| alert group id | 
 **name** | **string**| name of email groups | 

### Return type

[**EmailGroupsResp**](EmailGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmailGroup**
> EmailGroupResp UpdateEmailGroup(ctx, groupId, body)


update email group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| the id of email group | 
  **body** | [**EmailGroupUpdateReq**](EmailGroupUpdateReq.md)| email group | 

### Return type

[**EmailGroupResp**](EmailGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

