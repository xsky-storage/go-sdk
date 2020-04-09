# \FsClientGroupsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFSClients**](FsClientGroupsApi.md#AddFSClients) | **Put** /fs-client-groups/{fs_client_group_id}/fs-clients | 
[**CreateFSClientGroup**](FsClientGroupsApi.md#CreateFSClientGroup) | **Post** /fs-client-groups/ | 
[**DeleteFSClientGroup**](FsClientGroupsApi.md#DeleteFSClientGroup) | **Delete** /fs-client-groups/{fs_client_group_id} | 
[**GetFSClientGroup**](FsClientGroupsApi.md#GetFSClientGroup) | **Get** /fs-client-groups/{fs_client_group_id} | 
[**ListFSClientGroups**](FsClientGroupsApi.md#ListFSClientGroups) | **Get** /fs-client-groups/ | 
[**RemoveFSClients**](FsClientGroupsApi.md#RemoveFSClients) | **Delete** /fs-client-groups/{fs_client_group_id}/fs-clients | 
[**UpdateFSClientGroup**](FsClientGroupsApi.md#UpdateFSClientGroup) | **Patch** /fs-client-groups/{fs_client_group_id} | 


# **AddFSClients**
> FsClientGroupResp AddFSClients(ctx, fsClientGroupId, body)


add clients to file storage client group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsClientGroupId** | **int64**| client group id | 
  **body** | [**FsClientGroupAddClientsReq**](FsClientGroupAddClientsReq.md)| clients info | 

### Return type

[**FsClientGroupResp**](FSClientGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFSClientGroup**
> FsClientGroupResp CreateFSClientGroup(ctx, body)


Create file storage client group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsClientGroupCreateReq**](FsClientGroupCreateReq.md)| client group info | 

### Return type

[**FsClientGroupResp**](FSClientGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSClientGroup**
> DeleteFSClientGroup(ctx, fsClientGroupId)


delete file storage client group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsClientGroupId** | **int64**| client group id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSClientGroup**
> FsClientGroupResp GetFSClientGroup(ctx, fsClientGroupId)


Get file storage client group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsClientGroupId** | **int64**| client group id | 

### Return type

[**FsClientGroupResp**](FSClientGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSClientGroups**
> FsClientGroupsResp ListFSClientGroups(ctx, optional)


List file storage client groups

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
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 
 **fsClientId** | **int64**| file storage client id | 

### Return type

[**FsClientGroupsResp**](FSClientGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFSClients**
> FsClientGroupResp RemoveFSClients(ctx, fsClientGroupId, body)


remove clients from file storage client group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsClientGroupId** | **int64**| client group id | 
  **body** | [**FsClientGroupRemoveClientsReq**](FsClientGroupRemoveClientsReq.md)| clients info | 

### Return type

[**FsClientGroupResp**](FSClientGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSClientGroup**
> FsClientGroupResp UpdateFSClientGroup(ctx, fsClientGroupId, body)


Update file storage client group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsClientGroupId** | **int64**| client group id | 
  **body** | [**FsClientGroupUpdateReq**](FsClientGroupUpdateReq.md)| client group info | 

### Return type

[**FsClientGroupResp**](FSClientGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

