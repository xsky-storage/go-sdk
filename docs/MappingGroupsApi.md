# \MappingGroupsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVolumes**](MappingGroupsApi.md#AddVolumes) | **Post** /mapping-groups/{mapping_group_id}/block-volumes | 
[**CreateMappingGroup**](MappingGroupsApi.md#CreateMappingGroup) | **Post** /mapping-groups/ | 
[**DeleteMappingGroup**](MappingGroupsApi.md#DeleteMappingGroup) | **Delete** /mapping-groups/{mapping_group_id} | 
[**GetMappingGroup**](MappingGroupsApi.md#GetMappingGroup) | **Get** /mapping-groups/{mapping_group_id} | 
[**ListMappingGroups**](MappingGroupsApi.md#ListMappingGroups) | **Get** /mapping-groups/ | 
[**RemoveVolumes**](MappingGroupsApi.md#RemoveVolumes) | **Delete** /mapping-groups/{mapping_group_id}/block-volumes | 
[**UpdateMappingGroup**](MappingGroupsApi.md#UpdateMappingGroup) | **Patch** /mapping-groups/{mapping_group_id} | 
[**UpdateMappingGroupClientGroup**](MappingGroupsApi.md#UpdateMappingGroupClientGroup) | **Patch** /mapping-groups/{mapping_group_id}/client-group | 


# **AddVolumes**
> MappingGroupResp AddVolumes(ctx, mappingGroupId, blockVolumeIds)


add volumes to mapping group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingGroupId** | **int64**| mapping group id | 
  **blockVolumeIds** | [**MappingGroupAddVolumesReq**](MappingGroupAddVolumesReq.md)| block volume ids | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMappingGroup**
> MappingGroupResp CreateMappingGroup(ctx, mappingGroup)


create a mapping group in access path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingGroup** | [**MappingGroupCreateReq**](MappingGroupCreateReq.md)| mapping info | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMappingGroup**
> MappingGroupResp DeleteMappingGroup(ctx, mappingGroupId, optional)


Delete mapping group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingGroupId** | **int64**| mapping group id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingGroupId** | **int64**| mapping group id | 
 **force** | **bool**| force delete or not | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMappingGroup**
> MappingGroupResp GetMappingGroup(ctx, mappingGroupId)


Get mapping group by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingGroupId** | **int64**| mapping group id | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMappingGroups**
> MappingGroupsResp ListMappingGroups(ctx, optional)


List mapping groups

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
 **accessPathId** | **int64**| access path id | 
 **clientGroupId** | **int64**| client group id | 

### Return type

[**MappingGroupsResp**](MappingGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveVolumes**
> MappingGroupResp RemoveVolumes(ctx, mappingGroupId, blockVolumeIds, optional)


remove volumes from mapping group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingGroupId** | **int64**| mapping group id | 
  **blockVolumeIds** | [**MappingGroupRemoveVolumesReq**](MappingGroupRemoveVolumesReq.md)| block volume ids | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingGroupId** | **int64**| mapping group id | 
 **blockVolumeIds** | [**MappingGroupRemoveVolumesReq**](MappingGroupRemoveVolumesReq.md)| block volume ids | 
 **force** | **bool**| force delete or not | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMappingGroup**
> MappingGroupResp UpdateMappingGroup(ctx, mappingGroupId, mappingGroup)


update mapping group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingGroupId** | **int64**| mapping group id | 
  **mappingGroup** | [**MappingGroupUpdateReq**](MappingGroupUpdateReq.md)| mapping info | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMappingGroupClientGroup**
> MappingGroupResp UpdateMappingGroupClientGroup(ctx, mappingGroupId, clientGroupId)


update client group in mapping group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingGroupId** | **int64**| mapping group id | 
  **clientGroupId** | [**MappingGroupUpdateClientGroupReq**](MappingGroupUpdateClientGroupReq.md)| client group id | 

### Return type

[**MappingGroupResp**](MappingGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

