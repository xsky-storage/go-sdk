# \OsUsersApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectStorageUser**](OsUsersApi.md#CreateObjectStorageUser) | **Post** /os-users/ | 
[**DeleteObjectStorageUser**](OsUsersApi.md#DeleteObjectStorageUser) | **Delete** /os-users/{user_id} | 
[**GetObjectStorageUser**](OsUsersApi.md#GetObjectStorageUser) | **Get** /os-users/{user_id} | 
[**GetObjectStorageUserSamples**](OsUsersApi.md#GetObjectStorageUserSamples) | **Get** /os-users/{user_id}/samples | 
[**ListObjectStorageUsers**](OsUsersApi.md#ListObjectStorageUsers) | **Get** /os-users/ | 
[**UpdateObjectStorageUser**](OsUsersApi.md#UpdateObjectStorageUser) | **Patch** /os-users/{user_id} | 


# **CreateObjectStorageUser**
> ObjectStorageUserResp CreateObjectStorageUser(ctx, body)


create object storage user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ObjectStorageUserCreateReq**](ObjectStorageUserCreateReq.md)| user info | 

### Return type

[**ObjectStorageUserResp**](ObjectStorageUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteObjectStorageUser**
> ObjectStorageUserResp DeleteObjectStorageUser(ctx, userId)


delete object storage user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userId** | **int64**| user id | 

### Return type

[**ObjectStorageUserResp**](ObjectStorageUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectStorageUser**
> ObjectStorageUserResp GetObjectStorageUser(ctx, userId)


get object storage user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userId** | **int64**| user id | 

### Return type

[**ObjectStorageUserResp**](ObjectStorageUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectStorageUserSamples**
> ObjectStorageUserSamplesResp GetObjectStorageUserSamples(ctx, userId, optional)


get an object storage user's Samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userId** | **int64**| user id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64**| user id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**ObjectStorageUserSamplesResp**](ObjectStorageUserSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListObjectStorageUsers**
> ObjectStorageUsersResp ListObjectStorageUsers(ctx, optional)


List object storage users

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

### Return type

[**ObjectStorageUsersResp**](ObjectStorageUsersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateObjectStorageUser**
> ObjectStorageUserResp UpdateObjectStorageUser(ctx, userId, body)


update object storage user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userId** | **int64**| user id | 
  **body** | [**ObjectStorageUserUpdateReq**](ObjectStorageUserUpdateReq.md)| user info | 

### Return type

[**ObjectStorageUserResp**](ObjectStorageUserResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

