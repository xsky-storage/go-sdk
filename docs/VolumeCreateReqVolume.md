# VolumeCreateReqVolume

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSnapshotId** | **int64** | id of related block volume snapshot | [optional] [default to null]
**CrcCheck** | **bool** | crc check or not | [optional] [default to null]
**Description** | **string** | description of volume | [optional] [default to null]
**Flattened** | **bool** | flatten or not flatten | [optional] [default to null]
**Format** | **int64** | volume format: { 128 | 129 (advanced) }, default 128 | [optional] [default to null]
**Name** | **string** | name of volume | [default to null]
**PerformancePriority** | **int64** | performance priority: { 0 | 1 }, default 0 | [optional] [default to null]
**PoolId** | **int64** | id of pool belonged to | [default to null]
**Qos** | [***VolumeQosSpec**](VolumeQosSpec.md) | qos of volume | [optional] [default to null]
**QosEnabled** | **bool** | enable or disable the qos | [optional] [default to null]
**RemoteClusterFsId** | **string** | replication remote cluster fsid | [optional] [default to null]
**ReplicationPool** | **string** | replication peer pool | [optional] [default to null]
**ReplicationPoolId** | **int64** | replication peer pool id | [optional] [default to null]
**ReplicationPoolName** | **string** | replication peer pool name | [optional] [default to null]
**ReplicationVersion** | **int64** | replication version | [optional] [default to null]
**ReplicationVolume** | **string** | replication peer volume | [optional] [default to null]
**ReplicationVolumeId** | **int64** | replication peer volume id | [optional] [default to null]
**ReplicationVolumeName** | **string** | replication peer volume name | [optional] [default to null]
**Size** | **int64** | size of volume | [optional] [default to null]
**Sn** | **string** | volume sn, used when creating replication volume | [optional] [default to null]
**SnapshotReplicationPool** | **string** | snapshot replication peer pool | [optional] [default to null]
**SnapshotReplicationPoolId** | **int64** | snapshot replication peer pool id | [optional] [default to null]
**SnapshotReplicationVolume** | **string** | snapshot replication peer volume | [optional] [default to null]
**SnapshotReplicationVolumeId** | **int64** | snapshot replication peer volume id | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


