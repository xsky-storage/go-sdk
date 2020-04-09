# \ClusterApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BootNode**](ClusterApi.md#BootNode) | **Get** /cluster/bootnode | 
[**Cluster**](ClusterApi.md#Cluster) | **Get** /cluster/ | 
[**CommitMasterSwitching**](ClusterApi.md#CommitMasterSwitching) | **Post** /cluster/object-storage:commit-master-switching | 
[**DeleteObjectStorage**](ClusterApi.md#DeleteObjectStorage) | **Delete** /cluster/object-storage | 
[**EnableMultiZone**](ClusterApi.md#EnableMultiZone) | **Post** /cluster/object-storage:enable-multi-zone | 
[**GetActionLogReport**](ClusterApi.md#GetActionLogReport) | **Get** /cluster/action-log-report | 
[**GetAlertReport**](ClusterApi.md#GetAlertReport) | **Get** /cluster/alert-report | 
[**GetClusterOverview**](ClusterApi.md#GetClusterOverview) | **Get** /cluster/stats | 
[**GetClusterReport**](ClusterApi.md#GetClusterReport) | **Get** /cluster/report | 
[**GetClusterSamples**](ClusterApi.md#GetClusterSamples) | **Get** /cluster/samples | 
[**GetEventLogReport**](ClusterApi.md#GetEventLogReport) | **Get** /cluster/event-log-report | 
[**GetObjectStorage**](ClusterApi.md#GetObjectStorage) | **Get** /cluster/object-storage | 
[**GetStatsUsagePrediction**](ClusterApi.md#GetStatsUsagePrediction) | **Get** /cluster/stats-usage-prediction | 
[**InitObjectStorage**](ClusterApi.md#InitObjectStorage) | **Put** /cluster/object-storage | 
[**Installation**](ClusterApi.md#Installation) | **Get** /cluster/installation | 
[**PrepareMasterSwitching**](ClusterApi.md#PrepareMasterSwitching) | **Post** /cluster/object-storage:prepare-master-switching | 
[**RemoveClusterAccessInfo**](ClusterApi.md#RemoveClusterAccessInfo) | **Post** /cluster:remove-access-info | 
[**RollbackMasterSwitching**](ClusterApi.md#RollbackMasterSwitching) | **Post** /cluster/object-storage:rollback-master-switching | 
[**ServerTime**](ClusterApi.md#ServerTime) | **Get** /cluster/time | 
[**SetBootNode**](ClusterApi.md#SetBootNode) | **Put** /cluster/bootnode | 
[**SetClusterAccessInfo**](ClusterApi.md#SetClusterAccessInfo) | **Post** /cluster:set-access-info | 
[**SetObjectStorage**](ClusterApi.md#SetObjectStorage) | **Patch** /cluster/object-storage | 
[**SwitchOSZoneToMaster**](ClusterApi.md#SwitchOSZoneToMaster) | **Post** /cluster/object-storage:switch-os-zone-to-master | 
[**UpdateCluster**](ClusterApi.md#UpdateCluster) | **Patch** /cluster/ | 
[**UpdateClusterMaintenance**](ClusterApi.md#UpdateClusterMaintenance) | **Put** /cluster/maintenance | 


# **BootNode**
> BootNodeResp BootNode(ctx, )


get bootnode info

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BootNodeResp**](BootNodeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Cluster**
> ClusterRecordResp Cluster(ctx, )


Output the status of the cluster

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterRecordResp**](ClusterRecordResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommitMasterSwitching**
> ObjectStorageResp CommitMasterSwitching(ctx, )


Commit the switch to new master os zone.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ObjectStorageResp**](ObjectStorageResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteObjectStorage**
> ObjectStorageResp DeleteObjectStorage(ctx, optional)


Delete object storage system in current cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **bool**| force delete or not | 

### Return type

[**ObjectStorageResp**](ObjectStorageResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableMultiZone**
> ObjectStorageResp EnableMultiZone(ctx, )


Enable multi zone for object storage system

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ObjectStorageResp**](ObjectStorageResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActionLogReport**
> string GetActionLogReport(ctx, )


Get report of cluster action logs

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertReport**
> string GetAlertReport(ctx, )


Get report of cluster alerts

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterOverview**
> ClusterOverviewResp GetClusterOverview(ctx, )


Get overview data of a cluster

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterOverviewResp**](ClusterOverviewResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterReport**
> string GetClusterReport(ctx, )


Get report of a cluster

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterSamples**
> ClusterSamplesResp GetClusterSamples(ctx, optional)


get cluster's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**ClusterSamplesResp**](ClusterSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventLogReport**
> string GetEventLogReport(ctx, )


Get report of cluster event logs

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectStorage**
> ObjectStorageResp GetObjectStorage(ctx, )


Get object storage system status

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ObjectStorageResp**](ObjectStorageResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatsUsagePrediction**
> ClusterStatsPredictionResp GetStatsUsagePrediction(ctx, )


Get prediction of stats space usage

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterStatsPredictionResp**](ClusterStatsPredictionResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitObjectStorage**
> ObjectStorageResp InitObjectStorage(ctx, body)


Initialize object storage system in current cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ObjectStorageInitReq**](ObjectStorageInitReq.md)| init info | 

### Return type

[**ObjectStorageResp**](ObjectStorageResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Installation**
> ClusterInstallationResp Installation(ctx, )


show the installation status

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterInstallationResp**](ClusterInstallationResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrepareMasterSwitching**
> ObjectStorageResp PrepareMasterSwitching(ctx, )


Prepare to switch to new master os zone.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ObjectStorageResp**](ObjectStorageResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveClusterAccessInfo**
> ClusterResp RemoveClusterAccessInfo(ctx, )


Remove cluster access info

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterResp**](ClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RollbackMasterSwitching**
> ObjectStorageResp RollbackMasterSwitching(ctx, )


Rollback the switch to new master os zone.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ObjectStorageResp**](ObjectStorageResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServerTime**
> ClusterServerTimeResp ServerTime(ctx, )


show current server time

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterServerTimeResp**](ClusterServerTimeResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBootNode**
> BootNodeResp SetBootNode(ctx, body)


set bootnode info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**BootNodeReq**](BootNodeReq.md)| bootnode info | 

### Return type

[**BootNodeResp**](BootNodeResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetClusterAccessInfo**
> ClusterResp SetClusterAccessInfo(ctx, body)


Set cluster access info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ClusterSetAccessInfoReq**](ClusterSetAccessInfoReq.md)| access info | 

### Return type

[**ClusterResp**](ClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetObjectStorage**
> ObjectStorageResp SetObjectStorage(ctx, body)


Set object storage system in current cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ObjectStorageSetReq**](ObjectStorageSetReq.md)| set info | 

### Return type

[**ObjectStorageResp**](ObjectStorageResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchOSZoneToMaster**
> ObjectStorageResp SwitchOSZoneToMaster(ctx, optional)


Switch current os zone to master

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **bool**| force to switch even if old master is dead | 

### Return type

[**ObjectStorageResp**](ObjectStorageResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCluster**
> ClusterResp UpdateCluster(ctx, body)


set disk lighting mode of cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ClusterUpdateReq**](ClusterUpdateReq.md)| cluster disk lighting info | 

### Return type

[**ClusterResp**](ClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClusterMaintenance**
> ClusterResp UpdateClusterMaintenance(ctx, body)


update maintenance mode of cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ClusterMaintainReq**](ClusterMaintainReq.md)| cluster maintenance info | 

### Return type

[**ClusterResp**](ClusterResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

