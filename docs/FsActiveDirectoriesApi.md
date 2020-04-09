# \FsActiveDirectoriesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFSActiveDirectory**](FsActiveDirectoriesApi.md#CreateFSActiveDirectory) | **Post** /fs-active-directories/ | 
[**DeleteFSActiveDirectory**](FsActiveDirectoriesApi.md#DeleteFSActiveDirectory) | **Delete** /fs-active-directories/{fs_active_directory_id} | 
[**GetFSActiveDirectory**](FsActiveDirectoriesApi.md#GetFSActiveDirectory) | **Get** /fs-active-directories/{fs_active_directory_id} | 
[**ListFSActiveDirectories**](FsActiveDirectoriesApi.md#ListFSActiveDirectories) | **Get** /fs-active-directories/ | 
[**UpdateFSActiveDirectory**](FsActiveDirectoriesApi.md#UpdateFSActiveDirectory) | **Patch** /fs-active-directories/{fs_active_directory_id} | 
[**VerifyFSActiveDirectory**](FsActiveDirectoriesApi.md#VerifyFSActiveDirectory) | **Post** /fs-active-directories/{fs_active_directory_id}:verify | 


# **CreateFSActiveDirectory**
> FsActiveDirectoryResp CreateFSActiveDirectory(ctx, body)


create file storage active directory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsActiveDirectoryCreateReq**](FsActiveDirectoryCreateReq.md)| file storage active directory info | 

### Return type

[**FsActiveDirectoryResp**](FSActiveDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSActiveDirectory**
> FsActiveDirectoryResp DeleteFSActiveDirectory(ctx, fsActiveDirectoryId)


Delete file storage active directory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsActiveDirectoryId** | **int64**| file storage active directory id | 

### Return type

[**FsActiveDirectoryResp**](FSActiveDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSActiveDirectory**
> FsActiveDirectoryResp GetFSActiveDirectory(ctx, fsActiveDirectoryId)


Get a file storage active directory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsActiveDirectoryId** | **int64**| file storage active directory id | 

### Return type

[**FsActiveDirectoryResp**](FSActiveDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSActiveDirectories**
> FsActiveDirectoriesResp ListFSActiveDirectories(ctx, optional)


List file storage active directories

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

[**FsActiveDirectoriesResp**](FSActiveDirectoriesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSActiveDirectory**
> FsActiveDirectoryResp UpdateFSActiveDirectory(ctx, fsActiveDirectoryId, body)


Update a file storage active directory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsActiveDirectoryId** | **int64**| file storage active directory id | 
  **body** | [**FsActiveDirectoryUpdateReq**](FsActiveDirectoryUpdateReq.md)| file storage active directory info | 

### Return type

[**FsActiveDirectoryResp**](FSActiveDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyFSActiveDirectory**
> FsActiveDirectoryResp VerifyFSActiveDirectory(ctx, fsActiveDirectoryId)


Verify a fs active directory or user/group info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsActiveDirectoryId** | **int64**| file storage active directory id | 

### Return type

[**FsActiveDirectoryResp**](FSActiveDirectoryResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

