# \AlertsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountAlerts**](AlertsApi.md#CountAlerts) | **Get** /alerts/stats | 
[**DeleteAlert**](AlertsApi.md#DeleteAlert) | **Delete** /alerts/{alert_id} | 
[**DeleteAlerts**](AlertsApi.md#DeleteAlerts) | **Delete** /alerts/ | 
[**DoAlert**](AlertsApi.md#DoAlert) | **Put** /alerts/{alert_id} | 
[**DoAlerts**](AlertsApi.md#DoAlerts) | **Patch** /alerts/ | 
[**ListAlerts**](AlertsApi.md#ListAlerts) | **Get** /alerts/ | 
[**ResolveAlert**](AlertsApi.md#ResolveAlert) | **Post** /alerts/{alert_id}:resolve | 


# **CountAlerts**
> AlertStatsResp CountAlerts(ctx, optional)


count all alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acked** | **bool**| acked of alert | 
 **resolved** | **bool**| resolved or not of alert | 
 **resourceType** | **string**| resource type of alert | 
 **resourceId** | **int64**| resource id of alert | 

### Return type

[**AlertStatsResp**](AlertStatsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlert**
> DeleteAlert(ctx, alertId)


Delete alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **alertId** | **int64**| alert id | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlerts**
> DeleteAlerts(ctx, optional)


delete alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **string**| create time of last alert | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoAlert**
> AlertResp DoAlert(ctx, alertId, alert)


set the ack status of alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **alertId** | **int64**| the id of alert | 
  **alert** | [**AlertActionReq**](AlertActionReq.md)| the alert action info | 

### Return type

[**AlertResp**](AlertResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoAlerts**
> DoAlerts(ctx, alert)


set the ack status of alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **alert** | [**AlertsActionReq**](AlertsActionReq.md)| the alerts action info | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAlerts**
> AlertsResp ListAlerts(ctx, optional)


List all alerts

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
 **resourceType** | **string**| resource type of alert | 
 **resourceId** | **int64**| resource id of alert | 
 **alertType** | **string**| type of alert | 
 **acked** | **bool**| acked of alert | 
 **resolved** | **bool**| resolved or not of alert | 
 **resolveType** | **string**| resolve type of alert | 
 **level** | **string**| level of alert | 
 **durationBegin** | **string**| duration begin timestamp | 
 **durationEnd** | **string**| duration end timestamp | 

### Return type

[**AlertsResp**](AlertsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResolveAlert**
> AlertResp ResolveAlert(ctx, alertId)


set the resolved status of alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **alertId** | **int64**| the id of alert | 

### Return type

[**AlertResp**](AlertResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

