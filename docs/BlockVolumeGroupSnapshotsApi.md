# \BlockVolumeGroupSnapshotsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlockVolumeGroupSnapshot**](BlockVolumeGroupSnapshotsApi.md#CreateBlockVolumeGroupSnapshot) | **Post** /block-volume-group-snapshots/ | 
[**DeleteBlockVolumeGroupSnapshot**](BlockVolumeGroupSnapshotsApi.md#DeleteBlockVolumeGroupSnapshot) | **Delete** /block-volume-group-snapshots/{block_volume_group_snapshot_id} | 
[**GetBlockVolumeGroupSnapshot**](BlockVolumeGroupSnapshotsApi.md#GetBlockVolumeGroupSnapshot) | **Get** /block-volume-group-snapshots/{block_volume_group_snapshot_id} | 
[**ListBlockVolumeGroupSnapshots**](BlockVolumeGroupSnapshotsApi.md#ListBlockVolumeGroupSnapshots) | **Get** /block-volume-group-snapshots/ | 
[**UpdateBlockVolumeGroupSnapshot**](BlockVolumeGroupSnapshotsApi.md#UpdateBlockVolumeGroupSnapshot) | **Patch** /block-volume-group-snapshots/{block_volume_group_snapshot_id} | 


# **CreateBlockVolumeGroupSnapshot**
> VolumeGroupSnapshotResp CreateBlockVolumeGroupSnapshot(ctx, body)


Create block volume group snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**VolumeGroupSnapshotCreateReq**](VolumeGroupSnapshotCreateReq.md)| volume group snapshot info | 

### Return type

[**VolumeGroupSnapshotResp**](VolumeGroupSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBlockVolumeGroupSnapshot**
> VolumeGroupSnapshotResp DeleteBlockVolumeGroupSnapshot(ctx, blockVolumeGroupSnapshotId)


Delete a block volume group snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeGroupSnapshotId** | **int64**| block volume group snapshot id | 

### Return type

[**VolumeGroupSnapshotResp**](VolumeGroupSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockVolumeGroupSnapshot**
> SnapshotResp GetBlockVolumeGroupSnapshot(ctx, blockVolumeGroupSnapshotId)


get block volume group snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeGroupSnapshotId** | **int64**| the block volume group snapshot id | 

### Return type

[**SnapshotResp**](SnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBlockVolumeGroupSnapshots**
> VolumeGroupSnapshotsResp ListBlockVolumeGroupSnapshots(ctx, optional)


List block volume group snapshots

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

[**VolumeGroupSnapshotsResp**](VolumeGroupSnapshotsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBlockVolumeGroupSnapshot**
> VolumeGroupSnapshotResp UpdateBlockVolumeGroupSnapshot(ctx, blockVolumeGroupSnapshotId, body)


Update block volume group snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeGroupSnapshotId** | **int64**| the block volume group snapshot id | 
  **body** | [**VolumeGroupSnapshotUpdateReq**](VolumeGroupSnapshotUpdateReq.md)| volume group snapshot info | 

### Return type

[**VolumeGroupSnapshotResp**](VolumeGroupSnapshotResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

