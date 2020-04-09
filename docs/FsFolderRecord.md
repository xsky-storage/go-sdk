# FsFolderRecord

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPath** | [***AccessPathNestview**](AccessPath_Nestview.md) |  | [optional] [default to null]
**ActionStatus** | **string** |  | [optional] [default to null]
**BlockVolume** | [***VolumeNestview**](Volume_Nestview.md) |  | [optional] [default to null]
**BlockVolumeStatus** | **string** |  | [optional] [default to null]
**Create** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**DpFsSnapshotPolicy** | [***DpFsSnapshotPolicyNestview**](DpFSSnapshotPolicy_Nestview.md) |  | [optional] [default to null]
**Flattened** | **bool** |  | [optional] [default to null]
**Formatted** | **bool** |  | [optional] [default to null]
**FsGatewayGroup** | [***FsGatewayGroup**](FSGatewayGroup.md) |  | [optional] [default to null]
**FsQuotaTreeNum** | **int64** |  | [optional] [default to null]
**FsSnapshot** | [***FsSnapshotNestview**](FSSnapshot_Nestview.md) |  | [optional] [default to null]
**FsSnapshotNum** | **int64** |  | [optional] [default to null]
**Id** | **int64** |  | [optional] [default to null]
**LatestSnapshotTime** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Pool** | [***PoolNestview**](Pool_Nestview.md) |  | [optional] [default to null]
**Progress** | **float64** |  | [optional] [default to null]
**Qos** | [***VolumeQosSpec**](VolumeQosSpec.md) |  | [optional] [default to null]
**QosEnabled** | **bool** |  | [optional] [default to null]
**QuotaTreeShareTypes** | **[]string** |  | [optional] [default to null]
**QuotaTreeShared** | **bool** |  | [optional] [default to null]
**ShareTypes** | **[]string** |  | [optional] [default to null]
**Shared** | **bool** | before version 3.2.14, these fields was calculated by wizard, but there is quota trees in new verion, calculations is too complicated, so add the fields to folder struct | [optional] [default to null]
**Size** | **int64** |  | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Update** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Samples** | [**[]FsFolderStat**](FSFolderStat.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


