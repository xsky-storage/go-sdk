# \LicensesApi

All URIs are relative to *https://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadLicenseKey**](LicensesApi.md#DownloadLicenseKey) | **Get** /licenses/key | 
[**GetActiveLicense**](LicensesApi.md#GetActiveLicense) | **Get** /licenses/license | 
[**GetLicense**](LicensesApi.md#GetLicense) | **Get** /licenses/{license_id} | 
[**ListLicenses**](LicensesApi.md#ListLicenses) | **Get** /licenses/ | 
[**RegisterLicense**](LicensesApi.md#RegisterLicense) | **Post** /licenses/ | 


# **DownloadLicenseKey**
> string DownloadLicenseKey(ctx, )


get license key file

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActiveLicense**
> LicenseResp GetActiveLicense(ctx, )


get current active license

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LicenseResp**](LicenseResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLicense**
> LicenseResp GetLicense(ctx, licenseId)


get license

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **licenseId** | **int64**| the license id | 

### Return type

[**LicenseResp**](LicenseResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLicenses**
> LicensesResp ListLicenses(ctx, optional)


List licenses

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
 **active** | **bool**| filter license by active status | 

### Return type

[**LicensesResp**](LicensesResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterLicense**
> LicenseResp RegisterLicense(ctx, license, optional)


Activate product license

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **license** | ***os.File**| license info | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **license** | ***os.File**| license info | 
 **force** | **bool**| force activate | 

### Return type

[**LicenseResp**](LicenseResp.md)

### Authorization

[tokenInHeader](../README.md#tokenInHeader), [tokenInQuery](../README.md#tokenInQuery)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

