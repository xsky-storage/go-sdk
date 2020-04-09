# \BlockVolumeMigrationJobsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelBlockVolumeMigrationJob**](BlockVolumeMigrationJobsApi.md#CancelBlockVolumeMigrationJob) | **Post** /block-volume-migration-jobs/{block_volume_migration_job_id}:cancel | 
[**GetBlockVolumeMigrationJob**](BlockVolumeMigrationJobsApi.md#GetBlockVolumeMigrationJob) | **Get** /block-volume-migration-jobs/{block_volume_migration_job_id} | 
[**ListBlockVolumeMigrationJobs**](BlockVolumeMigrationJobsApi.md#ListBlockVolumeMigrationJobs) | **Get** /block-volume-migration-jobs/ | 
[**UpdateMigration**](BlockVolumeMigrationJobsApi.md#UpdateMigration) | **Post** /block-volume-migration-jobs/{block_volume_migration_job_id}:update | 


# **CancelBlockVolumeMigrationJob**
> BlockVolumeMigrationJobResp CancelBlockVolumeMigrationJob(ctx, blockVolumeMigrationJobId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeMigrationJobId** | **int64**| block volume migration job id | 

### Return type

[**BlockVolumeMigrationJobResp**](BlockVolumeMigrationJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockVolumeMigrationJob**
> BlockVolumeMigrationJobResp GetBlockVolumeMigrationJob(ctx, blockVolumeMigrationJobId)


get a volume migration job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeMigrationJobId** | **int64**| volume migration job id | 

### Return type

[**BlockVolumeMigrationJobResp**](BlockVolumeMigrationJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBlockVolumeMigrationJobs**
> BlockVolumeMigrationJobsResp ListBlockVolumeMigrationJobs(ctx, optional)


List volume migration jobs

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
 **status** | **string**| the status of volume migration job | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**BlockVolumeMigrationJobsResp**](BlockVolumeMigrationJobsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMigration**
> BlockVolumeMigrationJobResp UpdateMigration(ctx, blockVolumeMigrationJobId, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **blockVolumeMigrationJobId** | **int64**| block volume migration job id | 
  **body** | [**BlockVolumeUpdateMigrationReq**](BlockVolumeUpdateMigrationReq.md)| volume migration udpate info | 

### Return type

[**BlockVolumeMigrationJobResp**](BlockVolumeMigrationJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

