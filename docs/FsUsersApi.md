# \FsUsersApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFSUser**](FsUsersApi.md#CreateFSUser) | **Post** /fs-users/ | 
[**DeleteFSUser**](FsUsersApi.md#DeleteFSUser) | **Delete** /fs-users/{fs_user_id} | 
[**GetFSUser**](FsUsersApi.md#GetFSUser) | **Get** /fs-users/{fs_user_id} | 
[**ListFSUsers**](FsUsersApi.md#ListFSUsers) | **Get** /fs-users/ | 
[**UpdateFSUser**](FsUsersApi.md#UpdateFSUser) | **Patch** /fs-users/{fs_user_id} | 


# **CreateFSUser**
> FsUserResp CreateFSUser(ctx, body)


create file storage user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsUserCreateReq**](FsUserCreateReq.md)| user info | 

### Return type

[**FsUserResp**](FSUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSUser**
> DeleteFSUser(ctx, fsUserId)


delete file storage user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsUserId** | **int64**| user id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSUser**
> FsUserResp GetFSUser(ctx, fsUserId)


get file storage user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsUserId** | **int64**| user id | 

### Return type

[**FsUserResp**](FSUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSUsers**
> FsUsersResp ListFSUsers(ctx, optional)


List file storage users

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
 **security** | **string**| security of file storage users | 
 **fsUserGroupId** | **int64**| file storage user group id | 
 **notFsUserGroupId** | **int64**| file storage user group id | 
 **available** | **bool**| available user or not | 

### Return type

[**FsUsersResp**](FSUsersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSUser**
> FsUserResp UpdateFSUser(ctx, fsUserId, body)


update file storage user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsUserId** | **int64**| user id | 
  **body** | [**FsUserUpdateReq**](FsUserUpdateReq.md)| user info | 

### Return type

[**FsUserResp**](FSUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

