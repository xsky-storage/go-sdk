# \FsSnapshotsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFSSnapshot**](FsSnapshotsApi.md#CreateFSSnapshot) | **Post** /fs-snapshots/ | 
[**DeleteFSSnapshot**](FsSnapshotsApi.md#DeleteFSSnapshot) | **Delete** /fs-snapshots/{fs_snapshot_id} | 
[**GetFSSnapshot**](FsSnapshotsApi.md#GetFSSnapshot) | **Get** /fs-snapshots/{fs_snapshot_id} | 
[**ListFSSnapshots**](FsSnapshotsApi.md#ListFSSnapshots) | **Get** /fs-snapshots/ | 
[**UpdateFSSnapshot**](FsSnapshotsApi.md#UpdateFSSnapshot) | **Patch** /fs-snapshots/{fs_snapshot_id} | 


# **CreateFSSnapshot**
> FsSnapshotResp CreateFSSnapshot(ctx, body)


Create file storage snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**FsSnapshotCreateReq**](FsSnapshotCreateReq.md)| snapshot info | 

### Return type

[**FsSnapshotResp**](FSSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFSSnapshot**
> FsSnapshotResp DeleteFSSnapshot(ctx, fsSnapshotId, optional)


delete file storage snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsSnapshotId** | **int64**| snapshot id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fsSnapshotId** | **int64**| snapshot id | 
 **force** | **bool**| force delete or not | 

### Return type

[**FsSnapshotResp**](FSSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFSSnapshot**
> FsSnapshotResp GetFSSnapshot(ctx, fsSnapshotId)


Get file storage snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsSnapshotId** | **int64**| file storage snapshot id | 

### Return type

[**FsSnapshotResp**](FSSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFSSnapshots**
> FsSnapshotsResp ListFSSnapshots(ctx, optional)


List file storage snapshots

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
 **fsFolderId** | **int64**| file storage folder id | 

### Return type

[**FsSnapshotsResp**](FSSnapshotsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFSSnapshot**
> FsSnapshotResp UpdateFSSnapshot(ctx, fsSnapshotId, body)


Update file storage snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsSnapshotId** | **int64**| snapshot id | 
  **body** | [**FsSnapshotUpdateReq**](FsSnapshotUpdateReq.md)| snapshot info | 

### Return type

[**FsSnapshotResp**](FSSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

