# \BlockVolumesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchGetBlockVolumeSamples**](BlockVolumesApi.md#BatchGetBlockVolumeSamples) | **Get** /block-volumes/samples | 
[**CreateBlockVolume**](BlockVolumesApi.md#CreateBlockVolume) | **Post** /block-volumes/ | 
[**DeleteBlockVolume**](BlockVolumesApi.md#DeleteBlockVolume) | **Delete** /block-volumes/{block_volume_id} | 
[**GetBlockVolume**](BlockVolumesApi.md#GetBlockVolume) | **Get** /block-volumes/{block_volume_id} | 
[**GetBlockVolumeSamples**](BlockVolumesApi.md#GetBlockVolumeSamples) | **Get** /block-volumes/{block_volume_id}/samples | 
[**ListBlockVolumes**](BlockVolumesApi.md#ListBlockVolumes) | **Get** /block-volumes/ | 
[**MigrateBlockVolume**](BlockVolumesApi.md#MigrateBlockVolume) | **Post** /block-volumes/{block_volume_id}:migrate | 
[**RebuildBlockVolumeReplication**](BlockVolumesApi.md#RebuildBlockVolumeReplication) | **Post** /block-volumes/{block_volume_id}:rebuild-replication | 
[**SetBackupProtection**](BlockVolumesApi.md#SetBackupProtection) | **Post** /block-volumes/{block_volume_id}:set-backup-protection | 
[**SetBlockVolumeCrc**](BlockVolumesApi.md#SetBlockVolumeCrc) | **Post** /block-volumes/{block_volume_id}:set-crc | 
[**SetBlockVolumeReplication**](BlockVolumesApi.md#SetBlockVolumeReplication) | **Post** /block-volumes/{block_volume_id}:set-replication | 
[**SetBlockVolumeReplication_0**](BlockVolumesApi.md#SetBlockVolumeReplication_0) | **Post** /block-volumes/{block_volume_id}:unset-replication | 
[**SetSnapshotProtection**](BlockVolumesApi.md#SetSnapshotProtection) | **Post** /block-volumes/{block_volume_id}:set-snapshot-protection | 
[**SuspendBlockVolumeReplication**](BlockVolumesApi.md#SuspendBlockVolumeReplication) | **Post** /block-volumes/{block_volume_id}:suspend-replication | 
[**UnsetBackupProtection**](BlockVolumesApi.md#UnsetBackupProtection) | **Post** /block-volumes/{block_volume_id}:unset-backup-protection | 
[**UnsetBlockVolumeCrc**](BlockVolumesApi.md#UnsetBlockVolumeCrc) | **Post** /block-volumes/{block_volume_id}:unset-crc | 
[**UnsetSnapshotProtection**](BlockVolumesApi.md#UnsetSnapshotProtection) | **Post** /block-volumes/{block_volume_id}:unset-snapshot-protection | 
[**UpdateBlockVolume**](BlockVolumesApi.md#UpdateBlockVolume) | **Patch** /block-volumes/{block_volume_id} | 


# **BatchGetBlockVolumeSamples**
> MultiVolumesSamplesResp BatchGetBlockVolumeSamples(ctx, )


Get samples of multiple block volumes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MultiVolumesSamplesResp**](MultiVolumesSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBlockVolume**
> VolumeResp CreateBlockVolume(ctx, body)


Create block volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**VolumeCreateReq**](VolumeCreateReq.md)| volume info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBlockVolume**
> VolumeResp DeleteBlockVolume(ctx, blockVolumeId, optional)


Delete block volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| volume id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockVolumeId** | **int64**| volume id | 
 **bypassTrash** | **bool**| bypass trash or not | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockVolume**
> VolumeResp GetBlockVolume(ctx, blockVolumeId)


get a block volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| block volume id | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockVolumeSamples**
> VolumeSamplesResp GetBlockVolumeSamples(ctx, blockVolumeId, optional)


get a block volume's Samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| block volume id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockVolumeId** | **int64**| block volume id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**VolumeSamplesResp**](VolumeSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBlockVolumes**
> VolumesResp ListBlockVolumes(ctx, optional)


List block volumes

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
 **poolId** | **int64**| The id of the pool volumes belong to | 
 **blockSnapshotId** | **int64**| related snapshot id | 
 **name** | **string**| name of volume | 
 **clientGroupId** | **int64**| related client group id | 
 **mappingGroupId** | **int64**| related mapping group id | 
 **accessPathId** | **int64**| related access path id | 
 **passive** | **bool**| passive or not | 
 **recycled** | **bool**| recycled or not | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 
 **all** | **bool**| show all volumes | 
 **dpBlockBackupPolicyId** | **int64**| show volumes of specific dp block backup policy | 

### Return type

[**VolumesResp**](VolumesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateBlockVolume**
> VolumeResp MigrateBlockVolume(ctx, blockVolumeId, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| the block volume id | 
  **body** | [**VolumeMigrateReq**](VolumeMigrateReq.md)| volume migration info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebuildBlockVolumeReplication**
> VolumeResp RebuildBlockVolumeReplication(ctx, blockVolumeId, optional)


Rebuild block volume replication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| block volume id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockVolumeId** | **int64**| block volume id | 
 **force** | **bool**| force rebuild or not | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBackupProtection**
> VolumeResp SetBackupProtection(ctx, blockVolumeId, body)


Set backup protection for a block volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| the block volume id | 
  **body** | [**VolumeBackupProtectionReq**](VolumeBackupProtectionReq.md)| request info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBlockVolumeCrc**
> VolumeResp SetBlockVolumeCrc(ctx, blockVolumeId, body)


Set block volume crc

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| block volume id | 
  **body** | [**VolumeCrcSetReq**](VolumeCrcSetReq.md)| volume crc info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBlockVolumeReplication**
> VolumeResp SetBlockVolumeReplication(ctx, blockVolumeId, body)


Set block volume replication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| block volume id | 
  **body** | [**VolumeReplicationSetReq**](VolumeReplicationSetReq.md)| volume replication info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBlockVolumeReplication_0**
> VolumeResp SetBlockVolumeReplication_0(ctx, blockVolumeId)


Unset block volume replication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| block volume id | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSnapshotProtection**
> VolumeResp SetSnapshotProtection(ctx, blockVolumeId, body)


Set snapshot protection for a block volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| the block volume id | 
  **body** | [**VolumeSnapshotProtectionReq**](VolumeSnapshotProtectionReq.md)| request info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuspendBlockVolumeReplication**
> VolumeResp SuspendBlockVolumeReplication(ctx, blockVolumeId)


Suspend block volume replication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| block volume id | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsetBackupProtection**
> VolumeResp UnsetBackupProtection(ctx, blockVolumeId, optional)


Unset backup protection for a block volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| the block volume id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockVolumeId** | **int64**| the block volume id | 
 **force** | **bool**| force unset or not | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsetBlockVolumeCrc**
> VolumeResp UnsetBlockVolumeCrc(ctx, blockVolumeId, body)


Unset block volume crc

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| block volume id | 
  **body** | [**VolumeCrcSetReq**](VolumeCrcSetReq.md)| volume crc info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsetSnapshotProtection**
> VolumeResp UnsetSnapshotProtection(ctx, blockVolumeId, optional)


Unset snapshot protection for a block volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| the block volume id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockVolumeId** | **int64**| the block volume id | 
 **force** | **bool**| force unset or not | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBlockVolume**
> VolumeResp UpdateBlockVolume(ctx, blockVolumeId, body)


Update block volume info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeId** | **int64**| the block volume id | 
  **body** | [**VolumeUpdateReq**](VolumeUpdateReq.md)| volume info | 

### Return type

[**VolumeResp**](VolumeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

