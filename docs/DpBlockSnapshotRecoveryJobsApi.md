# \DpBlockSnapshotRecoveryJobsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpBlockSnapshotRecoveryJob**](DpBlockSnapshotRecoveryJobsApi.md#CreateDpBlockSnapshotRecoveryJob) | **Post** /dp-block-snapshot-recovery-jobs/ | 
[**DeleteDpBlockSnapshotRecoveryJob**](DpBlockSnapshotRecoveryJobsApi.md#DeleteDpBlockSnapshotRecoveryJob) | **Delete** /dp-block-snapshot-recovery-jobs/{job_id} | 
[**GetDpBlockSnapshotRecoveryJob**](DpBlockSnapshotRecoveryJobsApi.md#GetDpBlockSnapshotRecoveryJob) | **Get** /dp-block-snapshot-recovery-jobs/{job_id} | 
[**ListDpBlockSnapshotRecoveryJobs**](DpBlockSnapshotRecoveryJobsApi.md#ListDpBlockSnapshotRecoveryJobs) | **Get** /dp-block-snapshot-recovery-jobs/ | 


# **CreateDpBlockSnapshotRecoveryJob**
> DpBlockSnapshotRecoveryJobResp CreateDpBlockSnapshotRecoveryJob(ctx, body)


Create a dp block snapshot recovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**DpBlockSnapshotRecoveryJobCreateReq**](DpBlockSnapshotRecoveryJobCreateReq.md)| resource info | 

### Return type

[**DpBlockSnapshotRecoveryJobResp**](DpBlockSnapshotRecoveryJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDpBlockSnapshotRecoveryJob**
> DeleteDpBlockSnapshotRecoveryJob(ctx, jobId)


Delete a dp block snapshot recovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **int64**| resource id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDpBlockSnapshotRecoveryJob**
> DpBlockSnapshotRecoveryJobResp GetDpBlockSnapshotRecoveryJob(ctx, jobId)


Get a dp block snapshot recovery job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **int64**| resource id | 

### Return type

[**DpBlockSnapshotRecoveryJobResp**](DpBlockSnapshotRecoveryJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDpBlockSnapshotRecoveryJobs**
> DpBlockSnapshotRecoveryJobsResp ListDpBlockSnapshotRecoveryJobs(ctx, optional)


List dp block snapshot recovery jobs

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

[**DpBlockSnapshotRecoveryJobsResp**](DpBlockSnapshotRecoveryJobsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

