# \FsUserGroupsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFSUsers**](FsUserGroupsApi.md#AddFSUsers) | **Put** /fs-user-groups/{fs_user_group_id}/fs-users | 
[**CreateFSUserGroup**](FsUserGroupsApi.md#CreateFSUserGroup) | **Post** /fs-user-groups/ | 
[**DeleteFSUserGroup**](FsUserGroupsApi.md#DeleteFSUserGroup) | **Delete** /fs-user-groups/{fs_user_group_id} | 
[**GetFSUserGroup**](FsUserGroupsApi.md#GetFSUserGroup) | **Get** /fs-user-groups/{fs_user_group_id} | 
[**ListFSUserGroups**](FsUserGroupsApi.md#ListFSUserGroups) | **Get** /fs-user-groups/ | 
[**RemoveFSUsers**](FsUserGroupsApi.md#RemoveFSUsers) | **Delete** /fs-user-groups/{fs_user_group_id}/fs-users | 
[**UpdateFSUserGroup**](FsUserGroupsApi.md#UpdateFSUserGroup) | **Patch** /fs-user-groups/{fs_user_group_id} | 


# **AddFSUsers**
> FsUserGroupResp AddFSUsers(ctx, fsUserGroupId, body)


add users to file storage user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsUserGroupId** | **int64**| user group id | 
  **body** | [**FsUserGroupAddUsersReq**](FsUserGroupAddUsersReq.md)| users info | 

### Return type

[**FsUserGroupResp**](FSUserGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFSUserGroup**
> FsUserGroupResp CreateFSUserGroup(ctx, body)


Create file storage user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsUserGroupCreateReq**](FsUserGroupCreateReq.md)| user group info | 

### Return type

[**FsUserGroupResp**](FSUserGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSUserGroup**
> DeleteFSUserGroup(ctx, fsUserGroupId)


delete file storage user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsUserGroupId** | **int64**| user group id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSUserGroup**
> FsUserGroupResp GetFSUserGroup(ctx, fsUserGroupId)


Get file storage user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsUserGroupId** | **int64**| user group id | 

### Return type

[**FsUserGroupResp**](FSUserGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSUserGroups**
> FsUserGroupsResp ListFSUserGroups(ctx, optional)


List file storage user groups

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
 **type_** | **string**| security type of file storage user group | 
 **fsUserId** | **int64**| file storage user id | 

### Return type

[**FsUserGroupsResp**](FSUserGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFSUsers**
> FsUserGroupResp RemoveFSUsers(ctx, fsUserGroupId, body)


remove users from file storage user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsUserGroupId** | **int64**| user group id | 
  **body** | [**FsUserGroupRemoveUsersReq**](FsUserGroupRemoveUsersReq.md)| users info | 

### Return type

[**FsUserGroupResp**](FSUserGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSUserGroup**
> FsUserGroupResp UpdateFSUserGroup(ctx, fsUserGroupId, body)


Update file storage user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsUserGroupId** | **int64**| user group id | 
  **body** | [**FsUserGroupUpdateReq**](FsUserGroupUpdateReq.md)| user group info | 

### Return type

[**FsUserGroupResp**](FSUserGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

