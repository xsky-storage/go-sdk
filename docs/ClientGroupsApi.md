# \ClientGroupsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClientGroup**](ClientGroupsApi.md#CreateClientGroup) | **Post** /client-groups/ | 
[**DeleteClientGroup**](ClientGroupsApi.md#DeleteClientGroup) | **Delete** /client-groups/{client_group_id} | 
[**GetClientGroup**](ClientGroupsApi.md#GetClientGroup) | **Get** /client-groups/{client_group_id} | 
[**ListClientGroups**](ClientGroupsApi.md#ListClientGroups) | **Get** /client-groups/ | 
[**UpdateClientGroup**](ClientGroupsApi.md#UpdateClientGroup) | **Patch** /client-groups/{client_group_id} | 


# **CreateClientGroup**
> ClientGroupResp CreateClientGroup(ctx, client)


Create a client group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **client** | [**ClientGroupCreateReq**](ClientGroupCreateReq.md)| client group info | 

### Return type

[**ClientGroupResp**](ClientGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClientGroup**
> DeleteClientGroup(ctx, clientGroupId)


Delete a client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clientGroupId** | **int64**| client group id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientGroup**
> ClientGroupResp GetClientGroup(ctx, clientGroupId)


get a client group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clientGroupId** | **int64**| client group id | 

### Return type

[**ClientGroupResp**](ClientGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClientGroups**
> ClientGroupsResp ListClientGroups(ctx, optional)


List client groups

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
 **blockVolumeId** | **int64**| related block volume id | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**ClientGroupsResp**](ClientGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClientGroup**
> ClientGroupResp UpdateClientGroup(ctx, clientGroupId, client)


update client group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clientGroupId** | **int64**| client group id | 
  **client** | [**ClientGroupUpdateReq**](ClientGroupUpdateReq.md)| client group info | 

### Return type

[**ClientGroupResp**](ClientGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

