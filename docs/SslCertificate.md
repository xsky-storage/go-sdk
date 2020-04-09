# SslCertificate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | [**time.Time**](time.Time.md) | created time of certificate | [optional] [default to null]
**Description** | **string** | certificate description | [optional] [default to null]
**Enabled** | **bool** | enabled or not | [optional] [default to null]
**ForceHttps** | **bool** | redirect http request to https | [optional] [default to null]
**Id** | **int64** | certificate id | [optional] [default to null]
**Issuer** | [***interface{}**](interface{}.md) | issuer info | [optional] [default to null]
**Name** | **string** | certificate name | [optional] [default to null]
**NotAfter** | [**time.Time**](time.Time.md) | validity is not after the time | [optional] [default to null]
**NotBefore** | [**time.Time**](time.Time.md) | validity is not before the time | [optional] [default to null]
**PermittedDnsDomains** | [**[]interface{}**](interface{}.md) | permitted dns domains | [optional] [default to null]
**PublicKeyAlgorithm** | **string** | public key algorithm | [optional] [default to null]
**RawCertificate** | **string** | public certificate | [optional] [default to null]
**SignatureAlgorithm** | **string** | signature algorithm | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Subject** | [***interface{}**](interface{}.md) | subject info | [optional] [default to null]
**Type_** | **string** | applied type: admin, s3 | [optional] [default to null]
**Update** | [**time.Time**](time.Time.md) | updated time of certificate | [optional] [default to null]
**Version** | **int64** | certificate version | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


