# Go API client for client

This API allows users to interact through the MC API to provision and manage storage.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import client "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CopyVolumeDestinationPoolNameSourceGet**](docs/DefaultApi.md#copyvolumedestinationpoolnamesourceget) | **Get** /copy/volume/destination-pool/{destination-poolOption}/name/{nameOption}/{sourceOption} | 
*DefaultApi* | [**CreateSnapshotsVolumesNamesGet**](docs/DefaultApi.md#createsnapshotsvolumesnamesget) | **Get** /create/snapshots/volumes/{volumesOption}/{namesOption} | 
*DefaultApi* | [**CreateVolumePoolSizeNameGet**](docs/DefaultApi.md#createvolumepoolsizenameget) | **Get** /create/volume/pool/{poolOption}/size/{sizeOption}/{nameOption} | 
*DefaultApi* | [**CreateVolumePoolSizeTierAffinityNameGet**](docs/DefaultApi.md#createvolumepoolsizetieraffinitynameget) | **Get** /create/volume/pool/{poolOption}/size/{sizeOption}/tier-affinity/{tier-affinityOption}/{nameOption} | 
*DefaultApi* | [**DeleteHostsNamesGet**](docs/DefaultApi.md#deletehostsnamesget) | **Get** /delete/hosts/{namesOption} | 
*DefaultApi* | [**DeleteInitiatorNicknameNameGet**](docs/DefaultApi.md#deleteinitiatornicknamenameget) | **Get** /delete/initiator-nickname/{nameOption} | 
*DefaultApi* | [**DeleteSnapshotNamesGet**](docs/DefaultApi.md#deletesnapshotnamesget) | **Get** /delete/snapshot/{namesOption} | 
*DefaultApi* | [**DeleteVolumesNamesGet**](docs/DefaultApi.md#deletevolumesnamesget) | **Get** /delete/volumes/{namesOption} | 
*DefaultApi* | [**ExpandVolumeSizeNameGet**](docs/DefaultApi.md#expandvolumesizenameget) | **Get** /expand/volume/size/{sizeOption}/{nameOption} | 
*DefaultApi* | [**LoginGet**](docs/DefaultApi.md#loginget) | **Get** /login | 
*DefaultApi* | [**LogoutGet**](docs/DefaultApi.md#logoutget) | **Get** /logout | 
*DefaultApi* | [**MapVolumeAccessLunInitiatorNamesGet**](docs/DefaultApi.md#mapvolumeaccessluninitiatornamesget) | **Get** /map/volume/access/{accessOption}/lun/{lunOption}/initiator/{initiatorOption}/{namesOption} | 
*DefaultApi* | [**SchemaGet**](docs/DefaultApi.md#schemaget) | **Get** /meta/{schemaId} | 
*DefaultApi* | [**SetInitiatorIdNicknameGet**](docs/DefaultApi.md#setinitiatoridnicknameget) | **Get** /set/initiator/id/{idOption}/nickname/{nicknameOption} | 
*DefaultApi* | [**ShowAdvancedSettingsGet**](docs/DefaultApi.md#showadvancedsettingsget) | **Get** /show/advanced-settings | 
*DefaultApi* | [**ShowCacheParametersGet**](docs/DefaultApi.md#showcacheparametersget) | **Get** /show/cache-parameters | 
*DefaultApi* | [**ShowCertificateGet**](docs/DefaultApi.md#showcertificateget) | **Get** /show/certificate | 
*DefaultApi* | [**ShowControllerDateGet**](docs/DefaultApi.md#showcontrollerdateget) | **Get** /show/controller-date | 
*DefaultApi* | [**ShowControllersGet**](docs/DefaultApi.md#showcontrollersget) | **Get** /show/controllers | 
*DefaultApi* | [**ShowDiskGroupsGet**](docs/DefaultApi.md#showdiskgroupsget) | **Get** /show/disk-groups | 
*DefaultApi* | [**ShowDisksGet**](docs/DefaultApi.md#showdisksget) | **Get** /show/disks | 
*DefaultApi* | [**ShowEnclosuresGet**](docs/DefaultApi.md#showenclosuresget) | **Get** /show/enclosures | 
*DefaultApi* | [**ShowFansGet**](docs/DefaultApi.md#showfansget) | **Get** /show/fans | 
*DefaultApi* | [**ShowHostGroupsGet**](docs/DefaultApi.md#showhostgroupsget) | **Get** /show/host-groups | 
*DefaultApi* | [**ShowHostGroupsGroupsGet**](docs/DefaultApi.md#showhostgroupsgroupsget) | **Get** /show/host-groups/groups/{groupsOption} | 
*DefaultApi* | [**ShowHostGroupsHostsGet**](docs/DefaultApi.md#showhostgroupshostsget) | **Get** /show/host-groups/hosts/{hostsOption} | 
*DefaultApi* | [**ShowInitiatorNamesGet**](docs/DefaultApi.md#showinitiatornamesget) | **Get** /show/initiator/{namesOption} | 
*DefaultApi* | [**ShowInitiatorsGet**](docs/DefaultApi.md#showinitiatorsget) | **Get** /show/initiators | 
*DefaultApi* | [**ShowMapsAllGet**](docs/DefaultApi.md#showmapsallget) | **Get** /show/maps/all | 
*DefaultApi* | [**ShowMapsGet**](docs/DefaultApi.md#showmapsget) | **Get** /show/maps | 
*DefaultApi* | [**ShowMapsInitiatorGet**](docs/DefaultApi.md#showmapsinitiatorget) | **Get** /show/maps/initiator | 
*DefaultApi* | [**ShowMapsInitiatorNamesGet**](docs/DefaultApi.md#showmapsinitiatornamesget) | **Get** /show/maps/initiator/{namesOption} | 
*DefaultApi* | [**ShowMapsNamesGet**](docs/DefaultApi.md#showmapsnamesget) | **Get** /show/maps/{namesOption} | 
*DefaultApi* | [**ShowPoolsGet**](docs/DefaultApi.md#showpoolsget) | **Get** /show/pools | 
*DefaultApi* | [**ShowPowerSuppliesGet**](docs/DefaultApi.md#showpowersuppliesget) | **Get** /show/power-supplies | 
*DefaultApi* | [**ShowSnapshotsGet**](docs/DefaultApi.md#showsnapshotsget) | **Get** /show/snapshots | 
*DefaultApi* | [**ShowSnapshotsPatternGet**](docs/DefaultApi.md#showsnapshotspatternget) | **Get** /show/snapshots/pattern/{patternOption} | 
*DefaultApi* | [**ShowSnapshotsVolumeGet**](docs/DefaultApi.md#showsnapshotsvolumeget) | **Get** /show/snapshots/volume/{volumeOption} | 
*DefaultApi* | [**ShowSystemGet**](docs/DefaultApi.md#showsystemget) | **Get** /show/system | 
*DefaultApi* | [**ShowVersionsDetailGet**](docs/DefaultApi.md#showversionsdetailget) | **Get** /show/versions/detail | 
*DefaultApi* | [**ShowVersionsGet**](docs/DefaultApi.md#showversionsget) | **Get** /show/versions | 
*DefaultApi* | [**ShowVolumesNamesGet**](docs/DefaultApi.md#showvolumesnamesget) | **Get** /show/volumes/{namesOption} | 
*DefaultApi* | [**UnmapVolumeInitiatorNamesGet**](docs/DefaultApi.md#unmapvolumeinitiatornamesget) | **Get** /unmap/volume/initiator/{initiatorOption}/{namesOption} | 
*DefaultApi* | [**UnmapVolumeNamesGet**](docs/DefaultApi.md#unmapvolumenamesget) | **Get** /unmap/volume/{namesOption} | 


## Documentation For Models

 - [AdvancedSettingsTableObject](docs/AdvancedSettingsTableObject.md)
 - [AdvancedSettingsTableResourceInner](docs/AdvancedSettingsTableResourceInner.md)
 - [CacheSettingsObject](docs/CacheSettingsObject.md)
 - [CacheSettingsResourceInner](docs/CacheSettingsResourceInner.md)
 - [CertificateStatusObject](docs/CertificateStatusObject.md)
 - [CertificateStatusResourceInner](docs/CertificateStatusResourceInner.md)
 - [ControllerCacheParametersResourceInner](docs/ControllerCacheParametersResourceInner.md)
 - [ControllersObject](docs/ControllersObject.md)
 - [ControllersResourceInner](docs/ControllersResourceInner.md)
 - [DiskGroupsObject](docs/DiskGroupsObject.md)
 - [DiskGroupsResourceInner](docs/DiskGroupsResourceInner.md)
 - [DrivesObject](docs/DrivesObject.md)
 - [DrivesResourceInner](docs/DrivesResourceInner.md)
 - [EnclosuresObject](docs/EnclosuresObject.md)
 - [EnclosuresResourceInner](docs/EnclosuresResourceInner.md)
 - [ExpanderPortsResourceInner](docs/ExpanderPortsResourceInner.md)
 - [ExpandersResourceInner](docs/ExpandersResourceInner.md)
 - [FanObject](docs/FanObject.md)
 - [FanResourceInner](docs/FanResourceInner.md)
 - [FcPortResourceInner](docs/FcPortResourceInner.md)
 - [HostGroupObject](docs/HostGroupObject.md)
 - [HostGroupResourceInner](docs/HostGroupResourceInner.md)
 - [HostGroupViewObject](docs/HostGroupViewObject.md)
 - [HostGroupViewResourceInner](docs/HostGroupViewResourceInner.md)
 - [HostResourceInner](docs/HostResourceInner.md)
 - [HostViewMappingsResourceInner](docs/HostViewMappingsResourceInner.md)
 - [HostsViewObject](docs/HostsViewObject.md)
 - [HostsViewResourceInner](docs/HostsViewResourceInner.md)
 - [InitiatorObject](docs/InitiatorObject.md)
 - [InitiatorResourceInner](docs/InitiatorResourceInner.md)
 - [InitiatorViewObject](docs/InitiatorViewObject.md)
 - [InitiatorViewResourceInner](docs/InitiatorViewResourceInner.md)
 - [IscsiPortResourceInner](docs/IscsiPortResourceInner.md)
 - [NetworkParametersResourceInner](docs/NetworkParametersResourceInner.md)
 - [PoolsObject](docs/PoolsObject.md)
 - [PoolsResourceInner](docs/PoolsResourceInner.md)
 - [PortResourceInner](docs/PortResourceInner.md)
 - [PowerSuppliesObject](docs/PowerSuppliesObject.md)
 - [PowerSuppliesResourceInner](docs/PowerSuppliesResourceInner.md)
 - [RedundancyResourceInner](docs/RedundancyResourceInner.md)
 - [SasPortResourceInner](docs/SasPortResourceInner.md)
 - [ShowMapsInitiatorNamesGet200Response](docs/ShowMapsInitiatorNamesGet200Response.md)
 - [SnapshotsObject](docs/SnapshotsObject.md)
 - [SnapshotsResourceInner](docs/SnapshotsResourceInner.md)
 - [StatusObject](docs/StatusObject.md)
 - [StatusResourceInner](docs/StatusResourceInner.md)
 - [SystemObject](docs/SystemObject.md)
 - [SystemResourceInner](docs/SystemResourceInner.md)
 - [TiersResourceInner](docs/TiersResourceInner.md)
 - [VersionsObject](docs/VersionsObject.md)
 - [VersionsResourceInner](docs/VersionsResourceInner.md)
 - [VolumeGroupViewResourceInner](docs/VolumeGroupViewResourceInner.md)
 - [VolumeViewMappingsResourceInner](docs/VolumeViewMappingsResourceInner.md)
 - [VolumeViewObject](docs/VolumeViewObject.md)
 - [VolumeViewResourceInner](docs/VolumeViewResourceInner.md)
 - [VolumesObject](docs/VolumesObject.md)
 - [VolumesResourceInner](docs/VolumesResourceInner.md)


## Documentation For Authorization



### basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



