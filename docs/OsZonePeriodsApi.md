# \OsZonePeriodsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOSZonePeriod**](OsZonePeriodsApi.md#GetOSZonePeriod) | **Get** /os-zone-periods/{period_uuid} | 
[**ListOSZonePeriods**](OsZonePeriodsApi.md#ListOSZonePeriods) | **Get** /os-zone-periods/ | 


# **GetOSZonePeriod**
> OsZonePeriodResp GetOSZonePeriod(ctx, periodUuid)


Get a os zone period.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **periodUuid** | **string**| os zone period uuid | 

### Return type

[**OsZonePeriodResp**](OSZonePeriodResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOSZonePeriods**
> OsZonePeriodsResp ListOSZonePeriods(ctx, optional)


List os zone periods.

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

### Return type

[**OsZonePeriodsResp**](OSZonePeriodsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

