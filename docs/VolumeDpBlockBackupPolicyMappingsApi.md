# \VolumeDpBlockBackupPolicyMappingsApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVolumeDpBlockBackupPolicyMappings**](VolumeDpBlockBackupPolicyMappingsApi.md#ListVolumeDpBlockBackupPolicyMappings) | **Get** /volume-dp-block-backup-policy-mappings/ | 


# **ListVolumeDpBlockBackupPolicyMappings**
> VolumeDpBlockBackupPolicyMappingsResp ListVolumeDpBlockBackupPolicyMappings(ctx, optional)


List volume dp block backup policy mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockVolumeId** | **int64**| show mappings of specific block volume | 
 **dpBlockBackupPolicyId** | **int64**| show mappings of specific block volume | 

### Return type

[**VolumeDpBlockBackupPolicyMappingsResp**](VolumeDpBlockBackupPolicyMappingsResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

