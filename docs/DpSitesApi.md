# \DpSitesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDpSite**](DpSitesApi.md#CreateDpSite) | **Post** /dp-sites/ | 
[**DeleteDpSite**](DpSitesApi.md#DeleteDpSite) | **Delete** /dp-sites/{site_id} | 
[**GetBackupBlockSnapshots**](DpSitesApi.md#GetBackupBlockSnapshots) | **Get** /dp-sites/{site_id}/backup-block-snapshots | 
[**GetBackupBlockVolumes**](DpSitesApi.md#GetBackupBlockVolumes) | **Get** /dp-sites/{site_id}/backup-block-volumes | 
[**GetBackupClusters**](DpSitesApi.md#GetBackupClusters) | **Get** /dp-sites/{site_id}/backup-clusters | 
[**GetDpSite**](DpSitesApi.md#GetDpSite) | **Get** /dp-sites/{site_id} | 
[**ListDpSites**](DpSitesApi.md#ListDpSites) | **Get** /dp-sites/ | 
[**UpdateDpSite**](DpSitesApi.md#UpdateDpSite) | **Patch** /dp-sites/{site_id} | 


# **CreateDpSite**
> DpSiteResp CreateDpSite(ctx, body)


Create a data protection site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**DpSiteCreateReq**](DpSiteCreateReq.md)| dp site info | 

### Return type

[**DpSiteResp**](DpSiteResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDpSite**
> DeleteDpSite(ctx, siteId)


Delete data protection site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **siteId** | **int64**| dp site id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupBlockSnapshots**
> DpBackupBlockSnapshotsResp GetBackupBlockSnapshots(ctx, siteId, dpGatewayId, clusterFsId, blockVolumeName, optional)


List snapshots of a volume stored on a dp site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **siteId** | **int64**| dp site id | 
  **dpGatewayId** | **int64**| the dp gateway to execute the query | 
  **clusterFsId** | **string**| cluster fs id | 
  **blockVolumeName** | **string**| block volume name | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteId** | **int64**| dp site id | 
 **dpGatewayId** | **int64**| the dp gateway to execute the query | 
 **clusterFsId** | **string**| cluster fs id | 
 **blockVolumeName** | **string**| block volume name | 
 **after** | **string**| continuation token | 

### Return type

[**DpBackupBlockSnapshotsResp**](DpBackupBlockSnapshotsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupBlockVolumes**
> DpBackupBlockVolumesResp GetBackupBlockVolumes(ctx, siteId, dpGatewayId, clusterFsId, optional)


List volumes of a cluster stored on a dp site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **siteId** | **int64**| dp site id | 
  **dpGatewayId** | **int64**| the dp gateway to execute the query | 
  **clusterFsId** | **string**| cluster fs id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteId** | **int64**| dp site id | 
 **dpGatewayId** | **int64**| the dp gateway to execute the query | 
 **clusterFsId** | **string**| cluster fs id | 
 **after** | **string**| continuation token | 

### Return type

[**DpBackupBlockVolumesResp**](DpBackupBlockVolumesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupClusters**
> DpBackupClustersResp GetBackupClusters(ctx, siteId, dpGatewayId)


List clusters stored on a dp site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **siteId** | **int64**| dp site id | 
  **dpGatewayId** | **int64**| the dp gateway to execute the query | 

### Return type

[**DpBackupClustersResp**](DpBackupClustersResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDpSite**
> DpSiteResp GetDpSite(ctx, siteId)


Get data protection site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **siteId** | **int64**| protection site id | 

### Return type

[**DpSiteResp**](DpSiteResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDpSites**
> DpSitesResp ListDpSites(ctx, optional)


List data protection sites

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
 **protectionType** | **string**| protection type | 

### Return type

[**DpSitesResp**](DpSitesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDpSite**
> DpSiteResp UpdateDpSite(ctx, siteId, body)


Update a data protection site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **siteId** | **int64**| dp site id | 
  **body** | [**DpSiteUpdateReq**](DpSiteUpdateReq.md)| dp site info | 

### Return type

[**DpSiteResp**](DpSiteResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

