# \DpFsSnapshotJobsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDpFSSnapshotJob**](DpFsSnapshotJobsApi.md#GetDpFSSnapshotJob) | **Get** /dp-fs-snapshot-jobs/{job_id} | 
[**ListDpFSSnapshotJobs**](DpFsSnapshotJobsApi.md#ListDpFSSnapshotJobs) | **Get** /dp-fs-snapshot-jobs/ | 


# **GetDpFSSnapshotJob**
> DpFsSnapshotJobResp GetDpFSSnapshotJob(ctx, jobId)


Get dp fs snapshot job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobId** | **int64**| resource id | 

### Return type

[**DpFsSnapshotJobResp**](DpFSSnapshotJobResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDpFSSnapshotJobs**
> DpFsSnapshotJobsResp ListDpFSSnapshotJobs(ctx, optional)


List dp fs snapshot jobs

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
 **fsFolderId** | **int64**| fs folder | 
 **fsSnapshotId** | **int64**| fs snapshot | 
 **dpFsSnapshotPolicyId** | **int64**| dp fs snapshot policy | 

### Return type

[**DpFsSnapshotJobsResp**](DpFSSnapshotJobsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

