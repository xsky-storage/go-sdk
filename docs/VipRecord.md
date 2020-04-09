# VipRecord

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionStatus** | **string** |  | [optional] [default to null]
**Create** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Id** | **int64** |  | [optional] [default to null]
**Ip** | **string** |  | [optional] [default to null]
**MacAddress** | **string** |  | [optional] [default to null]
**Mask** | **int64** |  | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Update** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**VipGroup** | [***VipGroupNestview**](VIPGroup_Nestview.md) |  | [optional] [default to null]
**VirtualRouterId** | **int64** |  | [optional] [default to null]
**CurrentVipInstance** | [***VipInstanceNestview**](VIPInstance_Nestview.md) |  | [optional] [default to null]
**DefaultVipInstance** | [***VipInstanceNestview**](VIPInstance_Nestview.md) | Copy these two fields from corresponding fields of VIP to avoid JSON marshal recursion. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


