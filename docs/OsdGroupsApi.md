# \OsdGroupsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOsdsToOsdGroup**](OsdGroupsApi.md#AddOsdsToOsdGroup) | **Post** /osd-groups/{group_id}:add-osds | 
[**DisableDeviceTypeCheck**](OsdGroupsApi.md#DisableDeviceTypeCheck) | **Post** /osd-groups/{group_id}:disable-device-type-check | 
[**EnableDeviceTypeCheck**](OsdGroupsApi.md#EnableDeviceTypeCheck) | **Post** /osd-groups/{group_id}:enable-device-type-check | 
[**GetOsdGroup**](OsdGroupsApi.md#GetOsdGroup) | **Get** /osd-groups/{group_id} | 
[**GetOsdGroupSamples**](OsdGroupsApi.md#GetOsdGroupSamples) | **Get** /osd-groups/{group_id}/samples | 
[**ListOsdGroups**](OsdGroupsApi.md#ListOsdGroups) | **Get** /osd-groups/ | 
[**RemoveOsdsFromOsdGroup**](OsdGroupsApi.md#RemoveOsdsFromOsdGroup) | **Post** /osd-groups/{group_id}:remove-osds | 
[**ReweightOsdGroup**](OsdGroupsApi.md#ReweightOsdGroup) | **Post** /osd-groups/{group_id}:reweight | 
[**SetOsdFullRatio**](OsdGroupsApi.md#SetOsdFullRatio) | **Post** /osd-groups/{group_id}:set-osd-full-ratio | 
[**SetOsdGroupQos**](OsdGroupsApi.md#SetOsdGroupQos) | **Post** /osd-groups/{group_id}:set-qos | 


# **AddOsdsToOsdGroup**
> OsdGroupResp AddOsdsToOsdGroup(ctx, groupId, body)


Add osds to osd grouop

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| osd group id | 
  **body** | [**OsdGroupOsdsUpdateReq**](OsdGroupOsdsUpdateReq.md)| osd ids | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableDeviceTypeCheck**
> OsdGroupResp DisableDeviceTypeCheck(ctx, groupId)


Disable device type check when add osd

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| osd group id | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableDeviceTypeCheck**
> OsdGroupResp EnableDeviceTypeCheck(ctx, groupId)


Enable device type check when add osd

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| osd group id | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOsdGroup**
> OsdGroupResp GetOsdGroup(ctx, groupId)


Get osd group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| osd group id | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOsdGroupSamples**
> OsdGroupSamplesResp GetOsdGroupSamples(ctx, groupId, optional)


get a osd group's samples

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| osd group id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int64**| osd group id | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 
 **period** | **string**| samples period | 

### Return type

[**OsdGroupSamplesResp**](OsdGroupSamplesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOsdGroups**
> OsdGroupsResp ListOsdGroups(ctx, optional)


List osd groups

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

[**OsdGroupsResp**](OsdGroupsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOsdsFromOsdGroup**
> OsdGroupResp RemoveOsdsFromOsdGroup(ctx, groupId, body)


Remove multiple osds from a osd group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| osd group id | 
  **body** | [**OsdGroupOsdsUpdateReq**](OsdGroupOsdsUpdateReq.md)| remove osd ids | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReweightOsdGroup**
> OsdGroupResp ReweightOsdGroup(ctx, groupId)


Reweight pools of osd group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| osd group id | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetOsdFullRatio**
> OsdGroupResp SetOsdFullRatio(ctx, groupId, body)


Set osd full ratio of osd group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| osd group id | 
  **body** | [**SetOsdFullRatioReq**](SetOsdFullRatioReq.md)| osds full ratio | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetOsdGroupQos**
> OsdGroupResp SetOsdGroupQos(ctx, groupId, body)


Set osd group's qos

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupId** | **int64**| osd group id | 
  **body** | [**OsdGroupSetQosReq**](OsdGroupSetQosReq.md)| qos info | 

### Return type

[**OsdGroupResp**](OsdGroupResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

