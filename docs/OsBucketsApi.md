# \OsBucketsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomLabels**](OsBucketsApi.md#AddCustomLabels) | **Post** /os-buckets/{bucket_id}:add-custom-labels | 
[**AddOSReplicationPaths**](OsBucketsApi.md#AddOSReplicationPaths) | **Post** /os-buckets/{bucket_id}:add-os-replication-paths | 
[**AddOSReplicationZones**](OsBucketsApi.md#AddOSReplicationZones) | **Post** /os-buckets/{bucket_id}:add-os-replication-zones | 
[**BatchGetObjectStorageSamples**](OsBucketsApi.md#BatchGetObjectStorageSamples) | **Get** /os-buckets/samples | 
[**CreateBucket**](OsBucketsApi.md#CreateBucket) | **Post** /os-buckets/ | 
[**CreateObjectStorageBucketNFSClients**](OsBucketsApi.md#CreateObjectStorageBucketNFSClients) | **Post** /os-buckets/{bucket_id}/nfs-clients | 
[**DeleteBucket**](OsBucketsApi.md#DeleteBucket) | **Delete** /os-buckets/{bucket_id} | 
[**DeleteObjectStorageBucketNFSClients**](OsBucketsApi.md#DeleteObjectStorageBucketNFSClients) | **Delete** /os-buckets/{bucket_id}/nfs-clients | 
[**DeleteObjectStorageLifecycle**](OsBucketsApi.md#DeleteObjectStorageLifecycle) | **Delete** /os-buckets/{bucket_id}/lifecycle | 
[**GetBucket**](OsBucketsApi.md#GetBucket) | **Get** /os-buckets/{bucket_id} | 
[**GetObjectStorageBucketNFSClient**](OsBucketsApi.md#GetObjectStorageBucketNFSClient) | **Get** /os-buckets/{bucket_id}/nfs-clients/{client_id} | 
[**GetObjectStorageBucketSamples**](OsBucketsApi.md#GetObjectStorageBucketSamples) | **Get** /os-buckets/{bucket_id}/samples | 
[**ListBuckets**](OsBucketsApi.md#ListBuckets) | **Get** /os-buckets/ | 
[**ListObjectStorageBucketNFSClients**](OsBucketsApi.md#ListObjectStorageBucketNFSClients) | **Get** /os-buckets/{bucket_id}/nfs-clients | 
[**RemoveCustomLabels**](OsBucketsApi.md#RemoveCustomLabels) | **Post** /os-buckets/{bucket_id}:remove-custom-labels | 
[**RemoveOSBucketLoggings**](OsBucketsApi.md#RemoveOSBucketLoggings) | **Post** /os-buckets/{bucket_id}:remove-os-bucket-loggings | 
[**RemoveOSReplicationPaths**](OsBucketsApi.md#RemoveOSReplicationPaths) | **Post** /os-buckets/{bucket_id}:remove-os-replication-paths | 
[**RemoveOSReplicationZones**](OsBucketsApi.md#RemoveOSReplicationZones) | **Post** /os-buckets/{bucket_id}:remove-os-replication-zones | 
[**SetAccessLogging**](OsBucketsApi.md#SetAccessLogging) | **Post** /os-buckets/{bucket_id}:set-access-logging | 
[**SetMetadataSearch**](OsBucketsApi.md#SetMetadataSearch) | **Post** /os-buckets/{bucket_id}:set-metadata-search | 
[**SetOSBucketPolicy**](OsBucketsApi.md#SetOSBucketPolicy) | **Post** /os-buckets/{bucket_id}:set-bucket-policy | 
[**SetObjectStorageLifecycleRules**](OsBucketsApi.md#SetObjectStorageLifecycleRules) | **Put** /os-buckets/{bucket_id}/lifecycle | 
[**SuspendAccessLoggings**](OsBucketsApi.md#SuspendAccessLoggings) | **Post** /os-buckets/{bucket_id}:suspend-access-logging | 
[**SuspendOSReplicationPaths**](OsBucketsApi.md#SuspendOSReplicationPaths) | **Post** /os-buckets/{bucket_id}:suspend-os-replication-paths | 
[**SwitchOwnerOSZone**](OsBucketsApi.md#SwitchOwnerOSZone) | **Post** /os-buckets/{bucket_id}:switch-owner-os-zone | 
[**UnsetAccessLogging**](OsBucketsApi.md#UnsetAccessLogging) | **Post** /os-buckets/{bucket_id}:unset-access-logging | 
[**UnsetOSBucketPolicy**](OsBucketsApi.md#UnsetOSBucketPolicy) | **Post** /os-buckets/{bucket_id}:unset-bucket-policy | 
[**UnsuspendAccessLogging**](OsBucketsApi.md#UnsuspendAccessLogging) | **Post** /os-buckets/{bucket_id}:unsuspend-access-logging | 
[**UnsuspendOSReplicationPaths**](OsBucketsApi.md#UnsuspendOSReplicationPaths) | **Post** /os-buckets/{bucket_id}:unsuspend-os-replication-paths | 
[**UpdateBucket**](OsBucketsApi.md#UpdateBucket) | **Patch** /os-buckets/{bucket_id} | 
[**UpdateCustomLabels**](OsBucketsApi.md#UpdateCustomLabels) | **Post** /os-buckets/{bucket_id}:update-custom-labels | 
[**UpdateObjectStorageBucketNFSClient**](OsBucketsApi.md#UpdateObjectStorageBucketNFSClient) | **Patch** /os-buckets/{bucket_id}/nfs-clients/{client_id} | 


# **AddCustomLabels**
> ObjectStorageBucketResp AddCustomLabels(ctx, bucketId, body)


add object storage bucket custom labels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**OsBucketCustomLabelsAddReq**](OsBucketCustomLabelsAddReq.md)| bucket custom labels info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddOSReplicationPaths**
> ObjectStorageBucketResp AddOSReplicationPaths(ctx, bucketId, body)


add os replication paths for os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**OsBucketAddReplicationPathsReq**](OsBucketAddReplicationPathsReq.md)| replication paths info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddOSReplicationZones**
> ObjectStorageBucketResp AddOSReplicationZones(ctx, bucketId, body)


add os replication zones for os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**OsBucketAddReplicationZonesReq**](OsBucketAddReplicationZonesReq.md)| replication zones info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchGetObjectStorageSamples**
> MultiObjectStorageBucketsSamplesResp BatchGetObjectStorageSamples(ctx, ids)


Get samples of multiple object storage buckets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ids** | **string**| bucket ids | 

### Return type

[**MultiObjectStorageBucketsSamplesResp**](MultiObjectStorageBucketsSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBucket**
> ObjectStorageBucketResp CreateBucket(ctx, body)


Create os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ObjectStorageBucketCreateReq**](ObjectStorageBucketCreateReq.md)| bucket info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateObjectStorageBucketNFSClients**
> RawObjectStorageBucketResp CreateObjectStorageBucketNFSClients(ctx, bucketId, body)


create nfs client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**ObjectStorageBucketNfsClientsCreateReq**](ObjectStorageBucketNfsClientsCreateReq.md)| nfs client info | 

### Return type

[**RawObjectStorageBucketResp**](RawObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBucket**
> ObjectStorageBucketResp DeleteBucket(ctx, bucketId, optional)


delete object storage bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketId** | **int64**| bucket id | 
 **force** | **bool**| force delete or not | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteObjectStorageBucketNFSClients**
> RawObjectStorageBucketResp DeleteObjectStorageBucketNFSClients(ctx, bucketId, body)


delete nfs clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**ObjectStorageBucketNfsClientsDeleteReq**](ObjectStorageBucketNfsClientsDeleteReq.md)| nfs client info | 

### Return type

[**RawObjectStorageBucketResp**](RawObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteObjectStorageLifecycle**
> ObjectStorageBucketResp DeleteObjectStorageLifecycle(ctx, bucketId)


delete object storage lifecycle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBucket**
> ObjectStorageBucketResp GetBucket(ctx, bucketId)


Get object storage bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectStorageBucketNFSClient**
> ObjectStorageBucketNfsClientResp GetObjectStorageBucketNFSClient(ctx, bucketId, clientId)


show nfs client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **clientId** | **int64**| nfs client id | 

### Return type

[**ObjectStorageBucketNfsClientResp**](ObjectStorageBucketNFSClientResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectStorageBucketSamples**
> ObjectStorageBucketSamplesResp GetObjectStorageBucketSamples(ctx, bucketId, optional)


get an object storage bucket's Samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketId** | **int64**| bucket id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**ObjectStorageBucketSamplesResp**](ObjectStorageBucketSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBuckets**
> ObjectStorageBucketsResp ListBuckets(ctx, optional)


List object storage buckets

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
 **name** | **string**| name of object storage buckets | 
 **osPolicyId** | **int64**| The id of policy object storage buckets belong to | 
 **osUserId** | **int64**| The id of user object storage buckets belong to | 
 **replicationUuid** | **string**| The uuid of replication os buckets belong to | 
 **virtual** | **bool**| Virtual bucket or not | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**ObjectStorageBucketsResp**](ObjectStorageBucketsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListObjectStorageBucketNFSClients**
> ObjectStorageBucketNfsClientsResp ListObjectStorageBucketNFSClients(ctx, bucketId, optional)


List nfs clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketId** | **int64**| bucket id | 
 **limit** | **int64**| paging param | 
 **offset** | **int64**| paging param | 

### Return type

[**ObjectStorageBucketNfsClientsResp**](ObjectStorageBucketNFSClientsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveCustomLabels**
> ObjectStorageBucketResp RemoveCustomLabels(ctx, bucketId, body)


remove object storage bucket custom labels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| object storage bucket id | 
  **body** | [**OsBucketCustomLabelsRemoveReq**](OsBucketCustomLabelsRemoveReq.md)| custom labels info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOSBucketLoggings**
> ObjectStorageBucketResp RemoveOSBucketLoggings(ctx, bucketId, body)


Remove os bucket loggings of os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**OsBucketRemoveLoggingsReq**](OsBucketRemoveLoggingsReq.md)| os bucket loggings info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOSReplicationPaths**
> ObjectStorageBucketResp RemoveOSReplicationPaths(ctx, bucketId, body)


remove os replication paths for os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**OsBucketRemoveReplicationPathsReq**](OsBucketRemoveReplicationPathsReq.md)| replication paths info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOSReplicationZones**
> ObjectStorageBucketResp RemoveOSReplicationZones(ctx, bucketId, body, optional)


remove os replication zones for os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**OsBucketRemoveReplicationZonesReq**](OsBucketRemoveReplicationZonesReq.md)| replication zones info | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketId** | **int64**| bucket id | 
 **body** | [**OsBucketRemoveReplicationZonesReq**](OsBucketRemoveReplicationZonesReq.md)| replication zones info | 
 **force** | **bool**| force delete os replication zone even if target zone is dead | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetAccessLogging**
> ObjectStorageBucketResp SetAccessLogging(ctx, bucketId, body)


Set access logging of os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**OsBucketSetAccessLoggingReq**](OsBucketSetAccessLoggingReq.md)| access logging info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMetadataSearch**
> ObjectStorageBucketsResp SetMetadataSearch(ctx, bucketId, body)


set object storage bucket metadata search

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| object storage bucket id | 
  **body** | [**OsBucketMetadataSearchSetReq**](OsBucketMetadataSearchSetReq.md)| bucket metadata search info | 

### Return type

[**ObjectStorageBucketsResp**](ObjectStorageBucketsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetOSBucketPolicy**
> ObjectStorageBucketResp SetOSBucketPolicy(ctx, bucketId, body)


set object storage bucket policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**OsBucketPolicySetReq**](OsBucketPolicySetReq.md)| bucket policy info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetObjectStorageLifecycleRules**
> ObjectStorageBucketResp SetObjectStorageLifecycleRules(ctx, bucketId, body)


Set object storage lifecycle rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**ObjectStorageLifecycleSetReq**](ObjectStorageLifecycleSetReq.md)| lifecyce rules info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuspendAccessLoggings**
> ObjectStorageBucketResp SuspendAccessLoggings(ctx, bucketId)


suspend access logging for os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuspendOSReplicationPaths**
> ObjectStorageBucketResp SuspendOSReplicationPaths(ctx, bucketId, body)


suspend os replication paths for os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**OsBucketUpdateReplicationPathsReq**](OsBucketUpdateReplicationPathsReq.md)| replication paths info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchOwnerOSZone**
> ObjectStorageBucketResp SwitchOwnerOSZone(ctx, bucketId, body, optional)


Switch owner zone of os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**OsBucketSwitchOwnerOsZoneReq**](OsBucketSwitchOwnerOsZoneReq.md)| new owner os zone info | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketId** | **int64**| bucket id | 
 **body** | [**OsBucketSwitchOwnerOsZoneReq**](OsBucketSwitchOwnerOsZoneReq.md)| new owner os zone info | 
 **force** | **bool**| force to switch even if old owner zone is dead | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsetAccessLogging**
> ObjectStorageBucketResp UnsetAccessLogging(ctx, bucketId)


Unset access logging of os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsetOSBucketPolicy**
> ObjectStorageBucketResp UnsetOSBucketPolicy(ctx, bucketId)


unset object storage bucket policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsuspendAccessLogging**
> ObjectStorageBucketResp UnsuspendAccessLogging(ctx, bucketId)


unsuspend access logging for os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsuspendOSReplicationPaths**
> ObjectStorageBucketResp UnsuspendOSReplicationPaths(ctx, bucketId, body)


unsuspend os replication paths for os bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**OsBucketUpdateReplicationPathsReq**](OsBucketUpdateReplicationPathsReq.md)| replication paths info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBucket**
> ObjectStorageBucketResp UpdateBucket(ctx, bucketId, body)


Update object storage bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **body** | [**ObjectStorageBucketUpdateReq**](ObjectStorageBucketUpdateReq.md)| bucket info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomLabels**
> ObjectStorageBucketResp UpdateCustomLabels(ctx, bucketId, body)


update object storage bucket custom labels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| object storage bucket id | 
  **body** | [**OsBucketCustomLabelsUpdateReq**](OsBucketCustomLabelsUpdateReq.md)| custom labels info | 

### Return type

[**ObjectStorageBucketResp**](ObjectStorageBucketResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateObjectStorageBucketNFSClient**
> ObjectStorageBucketNfsClientResp UpdateObjectStorageBucketNFSClient(ctx, bucketId, clientId, body)


update nfs client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **bucketId** | **int64**| bucket id | 
  **clientId** | **int64**| nfs client id | 
  **body** | [**ObjectStorageBucketNfsClientUpdateReq**](ObjectStorageBucketNfsClientUpdateReq.md)| nfs client info | 

### Return type

[**ObjectStorageBucketNfsClientResp**](ObjectStorageBucketNFSClientResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

