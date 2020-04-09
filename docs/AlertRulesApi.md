# \AlertRulesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertRule**](AlertRulesApi.md#CreateAlertRule) | **Put** /alert-rules/ | 
[**CreateAlertRuleResourceBlacklist**](AlertRulesApi.md#CreateAlertRuleResourceBlacklist) | **Post** /alert-rules/{rule_id}/blacklist | 
[**CreateAlertRule_0**](AlertRulesApi.md#CreateAlertRule_0) | **Post** /alert-rules/ | 
[**DeleteAlertRule**](AlertRulesApi.md#DeleteAlertRule) | **Delete** /alert-rules/{rule_id} | 
[**DeleteAlertRuleResourceBlacklist**](AlertRulesApi.md#DeleteAlertRuleResourceBlacklist) | **Delete** /alert-rules/{rule_id}/blacklist | 
[**GetAlertRule**](AlertRulesApi.md#GetAlertRule) | **Get** /alert-rules/{rule_id} | 
[**GetAlertRuleSchema**](AlertRulesApi.md#GetAlertRuleSchema) | **Get** /alert-rules/schema | 
[**ListAlertRuleResourceBlacklist**](AlertRulesApi.md#ListAlertRuleResourceBlacklist) | **Get** /alert-rules/{rule_id}/blacklist | 
[**ListAlertRules**](AlertRulesApi.md#ListAlertRules) | **Get** /alert-rules/ | 
[**UpdateAlertRule**](AlertRulesApi.md#UpdateAlertRule) | **Patch** /alert-rules/{rule_id} | 


# **CreateAlertRule**
> AlertRuleResp CreateAlertRule(ctx, body)


create alert rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**AlertRuleCreateReq**](AlertRuleCreateReq.md)| alert rule | 

### Return type

[**AlertRuleResp**](AlertRuleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAlertRuleResourceBlacklist**
> AlertRuleResourceBlacklistResp CreateAlertRuleResourceBlacklist(ctx, ruleId, body)


create resource blacklist of alert rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ruleId** | **int64**| the id of alert rule | 
  **body** | [**UpdateAlertRuleResourceBlacklistReq**](UpdateAlertRuleResourceBlacklistReq.md)| resource blacklist | 

### Return type

[**AlertRuleResourceBlacklistResp**](AlertRuleResourceBlacklistResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAlertRule_0**
> AlertRuleResp CreateAlertRule_0(ctx, body)


create alert rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**AlertRuleCreateReq**](AlertRuleCreateReq.md)| alert rule | 

### Return type

[**AlertRuleResp**](AlertRuleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlertRule**
> DeleteAlertRule(ctx, ruleId)


delete alert rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ruleId** | **int64**| the id of alert rule | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlertRuleResourceBlacklist**
> DeleteAlertRuleResourceBlacklist(ctx, ruleId, body)


delete resource blacklist of alert rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ruleId** | **int64**| the id of alert rule | 
  **body** | [**UpdateAlertRuleResourceBlacklistReq**](UpdateAlertRuleResourceBlacklistReq.md)| resource blacklist | 

### Return type

 (empty response body)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertRule**
> AlertRuleResp GetAlertRule(ctx, ruleId)


get alert rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ruleId** | **int64**| the id of alert rule | 

### Return type

[**AlertRuleResp**](AlertRuleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertRuleSchema**
> AlertRuleSchemaResp GetAlertRuleSchema(ctx, )


get alert rule schema

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AlertRuleSchemaResp**](AlertRuleSchemaResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAlertRuleResourceBlacklist**
> AlertRuleResourceBlacklistResp ListAlertRuleResourceBlacklist(ctx, ruleId, blacklist, optional)


List all blacklist of alert rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ruleId** | **int64**| the id of alert rule | 
  **blacklist** | **string**| filter resource in blacklist or not | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleId** | **int64**| the id of alert rule | 
 **blacklist** | **string**| filter resource in blacklist or not | 
 **limit** | **int64**| paging param | 
 **offset** | **int64**| paging param | 

### Return type

[**AlertRuleResourceBlacklistResp**](AlertRuleResourceBlacklistResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAlertRules**
> AlertRulesResp ListAlertRules(ctx, optional)


List all alert rules

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
 **alertGroupId** | **int64**| alert group id | 
 **resourceType** | **string**| resource type of alert rule | 
 **level** | **string**| level of alert rule | 
 **enabled** | **bool**| enabled or not | 
 **q** | **string**| query param of search | 
 **sort** | **string**| sort param of search | 

### Return type

[**AlertRulesResp**](AlertRulesResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAlertRule**
> AlertRuleResp UpdateAlertRule(ctx, ruleId, body)


set alert rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ruleId** | **int64**| the id of alert rule | 
  **body** | [**AlertRuleUpdateReq**](AlertRuleUpdateReq.md)| alert rule | 

### Return type

[**AlertRuleResp**](AlertRuleResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

