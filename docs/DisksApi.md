# \DisksApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePartitions**](DisksApi.md#CreatePartitions) | **Put** /disks/{disk_id}/partitions | 
[**DeletePartitions**](DisksApi.md#DeletePartitions) | **Delete** /disks/{disk_id}/partitions | 
[**GetDisk**](DisksApi.md#GetDisk) | **Get** /disks/{disk_id} | 
[**GetDiskSamples**](DisksApi.md#GetDiskSamples) | **Get** /disks/{disk_id}/samples | 
[**ListDisks**](DisksApi.md#ListDisks) | **Get** /disks/ | 
[**UpdateDisk**](DisksApi.md#UpdateDisk) | **Patch** /disks/{disk_id} | 


# **CreatePartitions**
> DiskResp CreatePartitions(ctx, diskId, body, optional)


create cache partitions for disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **diskId** | **int64**| disk id | 
  **body** | [**PartitionsCreateReq**](PartitionsCreateReq.md)| partitions info | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diskId** | **int64**| disk id | 
 **body** | [**PartitionsCreateReq**](PartitionsCreateReq.md)| partitions info | 
 **num** | **int64**| num of partitions to create | 

### Return type

[**DiskResp**](DiskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePartitions**
> DiskResp DeletePartitions(ctx, diskId)


delete cache partitions of disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **diskId** | **int64**| disk id | 

### Return type

[**DiskResp**](DiskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDisk**
> DiskResp GetDisk(ctx, diskId)


Get a disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **diskId** | **int64**| disk id | 

### Return type

[**DiskResp**](DiskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiskSamples**
> DiskSamplesResp GetDiskSamples(ctx, diskId, optional)


get a disk's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **diskId** | **int64**| disk id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diskId** | **int64**| disk id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**DiskSamplesResp**](DiskSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDisks**
> DisksResp ListDisks(ctx, optional)


List all pyhsical disks in the system

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
 **hostId** | **int64**| host id | 
 **isCache** | **bool**| filter cache disk | 
 **used** | **bool**| filter used | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**DisksResp**](DisksResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDisk**
> DiskResp UpdateDisk(ctx, diskId, disk)


update disk info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **diskId** | **int64**| disk id | 
  **disk** | [**DiskUpdateReq**](DiskUpdateReq.md)| disk info | 

### Return type

[**DiskResp**](DiskResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

