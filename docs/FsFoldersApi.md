# \FsFoldersApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFSQuotaTrees**](FsFoldersApi.md#AddFSQuotaTrees) | **Put** /fs-folders/{fs_folder_id}/fs-quota-trees | 
[**CreateFolder**](FsFoldersApi.md#CreateFolder) | **Post** /fs-folders/ | 
[**DeleteFolder**](FsFoldersApi.md#DeleteFolder) | **Delete** /fs-folders/{fs_folder_id} | 
[**GetFSFolderSamples**](FsFoldersApi.md#GetFSFolderSamples) | **Get** /fs-folders/{fs_folder_id}/samples | 
[**GetFolder**](FsFoldersApi.md#GetFolder) | **Get** /fs-folders/{fs_folder_id} | 
[**ListFolders**](FsFoldersApi.md#ListFolders) | **Get** /fs-folders/ | 
[**RemoveFSQuotaTrees**](FsFoldersApi.md#RemoveFSQuotaTrees) | **Delete** /fs-folders/{fs_folder_id}/fs-quota-trees | 
[**RollBackFolder**](FsFoldersApi.md#RollBackFolder) | **Post** /fs-folders/{fs_folder_id}:roll-back | 
[**SetFSSnapshotProtection**](FsFoldersApi.md#SetFSSnapshotProtection) | **Post** /fs-folders/{fs_folder_id}:set-snapshot-protection | 
[**UnsetFSSnapshotProtection**](FsFoldersApi.md#UnsetFSSnapshotProtection) | **Post** /fs-folders/{fs_folder_id}:unset-snapshot-protection | 
[**UpdateFolder**](FsFoldersApi.md#UpdateFolder) | **Patch** /fs-folders/{fs_folder_id} | 


# **AddFSQuotaTrees**
> FsFolderResp AddFSQuotaTrees(ctx, fsFolderId, body)


add file storage quota trees

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFolderId** | **int64**| fs folder id | 
  **body** | [**FsFolderAddQuotaTreesReq**](FsFolderAddQuotaTreesReq.md)| quota trees info | 

### Return type

[**FsFolderResp**](FSFolderResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFolder**
> FsFolderResp CreateFolder(ctx, body)


Create file storage folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsFolderCreateReq**](FsFolderCreateReq.md)| folder info | 

### Return type

[**FsFolderResp**](FSFolderResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFolder**
> FsFolderResp DeleteFolder(ctx, fsFolderId, optional)


delete file storage folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFolderId** | **int64**| folder id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fsFolderId** | **int64**| folder id | 
 **force** | **bool**| force delete | 

### Return type

[**FsFolderResp**](FSFolderResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSFolderSamples**
> FsFolderSamplesResp GetFSFolderSamples(ctx, fsFolderId, optional)


get a file storage folder's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFolderId** | **int64**| file storage folder id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fsFolderId** | **int64**| file storage folder id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**FsFolderSamplesResp**](FSFolderSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolder**
> FsFolderResp GetFolder(ctx, fsFolderId)


Get file storage folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFolderId** | **int64**| folder id | 

### Return type

[**FsFolderResp**](FSFolderResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFolders**
> FsFoldersResp ListFolders(ctx, optional)


List file storage folders

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
 **fsUserId** | **int64**| file storage user id | 
 **fsUserGroupId** | **int64**| file storage user group id | 
 **fsGatewayGroupId** | **int64**| file storage gateway group id | 
 **fsSnapshotId** | **int64**| file storage snapshot id | 
 **fsClientId** | **int64**| file storage client id | 
 **fsClientGroupId** | **int64**| file storage client group id | 
 **poolId** | **int64**| pool id | 

### Return type

[**FsFoldersResp**](FSFoldersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFSQuotaTrees**
> FsFolderResp RemoveFSQuotaTrees(ctx, fsFolderId, body)


delete fs quota trees from fs folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFolderId** | **int64**| fs folder id | 
  **body** | [**FsFolderRemoveQuotaTreesReq**](FsFolderRemoveQuotaTreesReq.md)| quota trees info | 

### Return type

[**FsFolderResp**](FSFolderResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollBackFolder**
> FsFolderResp RollBackFolder(ctx, fsFolderId, body)


roll back file storage folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFolderId** | **int64**| folder id | 
  **body** | [**FsFolderRollBackReq**](FsFolderRollBackReq.md)| folder info | 

### Return type

[**FsFolderResp**](FSFolderResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetFSSnapshotProtection**
> FsFolderResp SetFSSnapshotProtection(ctx, fsFolderId, body)


Set snapshot protection for a fs folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFolderId** | **int64**| fs folder id in db | 
  **body** | [**FsFolderSnapshotProtectionReq**](FsFolderSnapshotProtectionReq.md)| request info | 

### Return type

[**FsFolderResp**](FSFolderResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsetFSSnapshotProtection**
> FsFolderResp UnsetFSSnapshotProtection(ctx, fsFolderId, optional)


Unset snapshot protection for a fs folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFolderId** | **int64**| fs_folder id in db | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fsFolderId** | **int64**| fs_folder id in db | 
 **force** | **bool**| force unset or not | 

### Return type

[**FsFolderResp**](FSFolderResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFolder**
> FsFolderResp UpdateFolder(ctx, fsFolderId, body)


Update file storage folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsFolderId** | **int64**| folder id | 
  **body** | [**FsFolderUpdateReq**](FsFolderUpdateReq.md)| folder info | 

### Return type

[**FsFolderResp**](FSFolderResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

