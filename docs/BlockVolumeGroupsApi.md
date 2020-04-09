# \BlockVolumeGroupsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBlockVolumeToBlockVolumeGroup**](BlockVolumeGroupsApi.md#AddBlockVolumeToBlockVolumeGroup) | **Put** /block-volume-groups/{block_volume_group_id}/block-volumes | 
[**CreateBlockVolumeGroup**](BlockVolumeGroupsApi.md#CreateBlockVolumeGroup) | **Post** /block-volume-groups/ | 
[**DeleteBlockVolumeGroup**](BlockVolumeGroupsApi.md#DeleteBlockVolumeGroup) | **Delete** /block-volume-groups/{block_volume_group_id} | 
[**GetBlockVolumeGroup**](BlockVolumeGroupsApi.md#GetBlockVolumeGroup) | **Get** /block-volume-groups/{block_volume_group_id} | 
[**ListBlockVolumeGroups**](BlockVolumeGroupsApi.md#ListBlockVolumeGroups) | **Get** /block-volume-groups/ | 
[**RemoveBlockVolumeFromBlockVolumeGroup**](BlockVolumeGroupsApi.md#RemoveBlockVolumeFromBlockVolumeGroup) | **Delete** /block-volume-groups/{block_volume_group_id}/block-volumes | 
[**UpdateBlockVolumeGroup**](BlockVolumeGroupsApi.md#UpdateBlockVolumeGroup) | **Patch** /block-volume-groups/{block_volume_group_id} | 


# **AddBlockVolumeToBlockVolumeGroup**
> VolumeGroupResp AddBlockVolumeToBlockVolumeGroup(ctx, blockVolumeGroupId, body)


Add block volume to block volume group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeGroupId** | **int64**| block volume group id | 
  **body** | [**VolumeGroupAddVolumeReq**](VolumeGroupAddVolumeReq.md)| volume ids | 

### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBlockVolumeGroup**
> VolumeGroupResp CreateBlockVolumeGroup(ctx, body)


Create block volume group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**VolumeGroupCreateReq**](VolumeGroupCreateReq.md)| volume group info | 

### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBlockVolumeGroup**
> DeleteBlockVolumeGroup(ctx, blockVolumeGroupId)


Delete a block volume group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeGroupId** | **int64**| block volume group id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockVolumeGroup**
> VolumeGroupResp GetBlockVolumeGroup(ctx, blockVolumeGroupId)


get a block volume group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeGroupId** | **int64**| block volume group id | 

### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBlockVolumeGroups**
> VolumeGroupsResp ListBlockVolumeGroups(ctx, optional)


List block volume groups

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

[**VolumeGroupsResp**](VolumeGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveBlockVolumeFromBlockVolumeGroup**
> VolumeGroupResp RemoveBlockVolumeFromBlockVolumeGroup(ctx, blockVolumeGroupId, body)


Remove block volume from block volume group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeGroupId** | **int64**| block volume group id | 
  **body** | [**VolumeGroupRemoveVolumeReq**](VolumeGroupRemoveVolumeReq.md)| volume ids | 

### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBlockVolumeGroup**
> VolumeGroupResp UpdateBlockVolumeGroup(ctx, blockVolumeGroupId, body)


Update block volume group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeGroupId** | **int64**| block volume group id | 
  **body** | [**VolumeGroupUpdateReq**](VolumeGroupUpdateReq.md)| volume group info | 

### Return type

[**VolumeGroupResp**](VolumeGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

