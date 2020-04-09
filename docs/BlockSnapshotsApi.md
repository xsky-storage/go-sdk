# \BlockSnapshotsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlockSnapshot**](BlockSnapshotsApi.md#CreateBlockSnapshot) | **Post** /block-snapshots/ | 
[**DeleteBlockSnapshot**](BlockSnapshotsApi.md#DeleteBlockSnapshot) | **Delete** /block-snapshots/{snapshot_id} | 
[**GetBlockSnapshot**](BlockSnapshotsApi.md#GetBlockSnapshot) | **Get** /block-snapshots/{block_snapshot_id} | 
[**ListBlockSnapshots**](BlockSnapshotsApi.md#ListBlockSnapshots) | **Get** /block-snapshots/ | 
[**UpdateSnapshot**](BlockSnapshotsApi.md#UpdateSnapshot) | **Patch** /block-snapshots/{snapshot_id} | 


# **CreateBlockSnapshot**
> SnapshotResp CreateBlockSnapshot(ctx, body)


Create block snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**SnapshotCreateReq**](SnapshotCreateReq.md)| snapshot info | 

### Return type

[**SnapshotResp**](SnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBlockSnapshot**
> SnapshotResp DeleteBlockSnapshot(ctx, snapshotId)


Delete block snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotId** | **int64**| snapshot id | 

### Return type

[**SnapshotResp**](SnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockSnapshot**
> SnapshotResp GetBlockSnapshot(ctx, blockSnapshotId)


get block snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockSnapshotId** | **int64**| the block snapshot id | 

### Return type

[**SnapshotResp**](SnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBlockSnapshots**
> SnapshotsResp ListBlockSnapshots(ctx, optional)


List block snapshots

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
 **poolId** | **int64**| pool id | 
 **blockVolumeId** | **int64**| volume id | 
 **reserved** | **bool**| show reserved snapshot or not, default not | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 
 **all** | **bool**| show all snapshots | 

### Return type

[**SnapshotsResp**](SnapshotsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSnapshot**
> SnapshotResp UpdateSnapshot(ctx, snapshotId, body)


Update block snapshot info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotId** | **int64**| snapshot id | 
  **body** | [**SnapshotUpdateReq**](SnapshotUpdateReq.md)| snapshot info | 

### Return type

[**SnapshotResp**](SnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

