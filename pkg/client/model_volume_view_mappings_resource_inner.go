/*
Seagate Management Controller (MC) OpenAPI

This API allows users to interact through the MC API to provision and manage storage.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the VolumeViewMappingsResourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeViewMappingsResourceInner{}

// VolumeViewMappingsResourceInner struct for VolumeViewMappingsResourceInner
type VolumeViewMappingsResourceInner struct {
	ObjectName *string `json:"object-name,omitempty"`
	Meta       *string `json:"meta,omitempty"`
	// Access rights for this volume
	Access *string `json:"access,omitempty"`
	// Access rights for this volume( In numeric form )
	AccessNumeric      *int64  `json:"access-numeric,omitempty"`
	DurableId          *string `json:"durable-id,omitempty"`
	HostProfile        *string `json:"host-profile,omitempty"`
	HostProfileNumeric *int64  `json:"host-profile-numeric,omitempty"`
	// WWPN or IQN or Host Serial Number or Host Group Serial Number
	Identifier    *string `json:"identifier,omitempty"`
	InitiatorsUrl *string `json:"initiators-url,omitempty"`
	// Logical Unit Number
	Lun      *string `json:"lun,omitempty"`
	MappedId *string `json:"mapped-id,omitempty"`
	// User-defined alias for the mapped host
	Nickname *string `json:"nickname,omitempty"`
	ParentId *string `json:"parent-id,omitempty"`
	Ports    *string `json:"ports,omitempty"`
}

// NewVolumeViewMappingsResourceInner instantiates a new VolumeViewMappingsResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeViewMappingsResourceInner() *VolumeViewMappingsResourceInner {
	this := VolumeViewMappingsResourceInner{}
	return &this
}

// NewVolumeViewMappingsResourceInnerWithDefaults instantiates a new VolumeViewMappingsResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeViewMappingsResourceInnerWithDefaults() *VolumeViewMappingsResourceInner {
	this := VolumeViewMappingsResourceInner{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *VolumeViewMappingsResourceInner) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetMeta() string {
	if o == nil || IsNil(o.Meta) {
		var ret string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetMetaOk() (*string, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given string and assigns it to the Meta field.
func (o *VolumeViewMappingsResourceInner) SetMeta(v string) {
	o.Meta = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *VolumeViewMappingsResourceInner) SetAccess(v string) {
	o.Access = &v
}

// GetAccessNumeric returns the AccessNumeric field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetAccessNumeric() int64 {
	if o == nil || IsNil(o.AccessNumeric) {
		var ret int64
		return ret
	}
	return *o.AccessNumeric
}

// GetAccessNumericOk returns a tuple with the AccessNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetAccessNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessNumeric) {
		return nil, false
	}
	return o.AccessNumeric, true
}

// HasAccessNumeric returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasAccessNumeric() bool {
	if o != nil && !IsNil(o.AccessNumeric) {
		return true
	}

	return false
}

// SetAccessNumeric gets a reference to the given int64 and assigns it to the AccessNumeric field.
func (o *VolumeViewMappingsResourceInner) SetAccessNumeric(v int64) {
	o.AccessNumeric = &v
}

// GetDurableId returns the DurableId field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetDurableId() string {
	if o == nil || IsNil(o.DurableId) {
		var ret string
		return ret
	}
	return *o.DurableId
}

// GetDurableIdOk returns a tuple with the DurableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetDurableIdOk() (*string, bool) {
	if o == nil || IsNil(o.DurableId) {
		return nil, false
	}
	return o.DurableId, true
}

// HasDurableId returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasDurableId() bool {
	if o != nil && !IsNil(o.DurableId) {
		return true
	}

	return false
}

// SetDurableId gets a reference to the given string and assigns it to the DurableId field.
func (o *VolumeViewMappingsResourceInner) SetDurableId(v string) {
	o.DurableId = &v
}

// GetHostProfile returns the HostProfile field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetHostProfile() string {
	if o == nil || IsNil(o.HostProfile) {
		var ret string
		return ret
	}
	return *o.HostProfile
}

// GetHostProfileOk returns a tuple with the HostProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetHostProfileOk() (*string, bool) {
	if o == nil || IsNil(o.HostProfile) {
		return nil, false
	}
	return o.HostProfile, true
}

// HasHostProfile returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasHostProfile() bool {
	if o != nil && !IsNil(o.HostProfile) {
		return true
	}

	return false
}

// SetHostProfile gets a reference to the given string and assigns it to the HostProfile field.
func (o *VolumeViewMappingsResourceInner) SetHostProfile(v string) {
	o.HostProfile = &v
}

// GetHostProfileNumeric returns the HostProfileNumeric field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetHostProfileNumeric() int64 {
	if o == nil || IsNil(o.HostProfileNumeric) {
		var ret int64
		return ret
	}
	return *o.HostProfileNumeric
}

// GetHostProfileNumericOk returns a tuple with the HostProfileNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetHostProfileNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.HostProfileNumeric) {
		return nil, false
	}
	return o.HostProfileNumeric, true
}

// HasHostProfileNumeric returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasHostProfileNumeric() bool {
	if o != nil && !IsNil(o.HostProfileNumeric) {
		return true
	}

	return false
}

// SetHostProfileNumeric gets a reference to the given int64 and assigns it to the HostProfileNumeric field.
func (o *VolumeViewMappingsResourceInner) SetHostProfileNumeric(v int64) {
	o.HostProfileNumeric = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *VolumeViewMappingsResourceInner) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetInitiatorsUrl returns the InitiatorsUrl field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetInitiatorsUrl() string {
	if o == nil || IsNil(o.InitiatorsUrl) {
		var ret string
		return ret
	}
	return *o.InitiatorsUrl
}

// GetInitiatorsUrlOk returns a tuple with the InitiatorsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetInitiatorsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatorsUrl) {
		return nil, false
	}
	return o.InitiatorsUrl, true
}

// HasInitiatorsUrl returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasInitiatorsUrl() bool {
	if o != nil && !IsNil(o.InitiatorsUrl) {
		return true
	}

	return false
}

// SetInitiatorsUrl gets a reference to the given string and assigns it to the InitiatorsUrl field.
func (o *VolumeViewMappingsResourceInner) SetInitiatorsUrl(v string) {
	o.InitiatorsUrl = &v
}

// GetLun returns the Lun field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetLun() string {
	if o == nil || IsNil(o.Lun) {
		var ret string
		return ret
	}
	return *o.Lun
}

// GetLunOk returns a tuple with the Lun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetLunOk() (*string, bool) {
	if o == nil || IsNil(o.Lun) {
		return nil, false
	}
	return o.Lun, true
}

// HasLun returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasLun() bool {
	if o != nil && !IsNil(o.Lun) {
		return true
	}

	return false
}

// SetLun gets a reference to the given string and assigns it to the Lun field.
func (o *VolumeViewMappingsResourceInner) SetLun(v string) {
	o.Lun = &v
}

// GetMappedId returns the MappedId field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetMappedId() string {
	if o == nil || IsNil(o.MappedId) {
		var ret string
		return ret
	}
	return *o.MappedId
}

// GetMappedIdOk returns a tuple with the MappedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetMappedIdOk() (*string, bool) {
	if o == nil || IsNil(o.MappedId) {
		return nil, false
	}
	return o.MappedId, true
}

// HasMappedId returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasMappedId() bool {
	if o != nil && !IsNil(o.MappedId) {
		return true
	}

	return false
}

// SetMappedId gets a reference to the given string and assigns it to the MappedId field.
func (o *VolumeViewMappingsResourceInner) SetMappedId(v string) {
	o.MappedId = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetNickname() string {
	if o == nil || IsNil(o.Nickname) {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname) {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasNickname() bool {
	if o != nil && !IsNil(o.Nickname) {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *VolumeViewMappingsResourceInner) SetNickname(v string) {
	o.Nickname = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *VolumeViewMappingsResourceInner) SetParentId(v string) {
	o.ParentId = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *VolumeViewMappingsResourceInner) GetPorts() string {
	if o == nil || IsNil(o.Ports) {
		var ret string
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewMappingsResourceInner) GetPortsOk() (*string, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *VolumeViewMappingsResourceInner) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given string and assigns it to the Ports field.
func (o *VolumeViewMappingsResourceInner) SetPorts(v string) {
	o.Ports = &v
}

func (o VolumeViewMappingsResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeViewMappingsResourceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectName) {
		toSerialize["object-name"] = o.ObjectName
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.AccessNumeric) {
		toSerialize["access-numeric"] = o.AccessNumeric
	}
	if !IsNil(o.DurableId) {
		toSerialize["durable-id"] = o.DurableId
	}
	if !IsNil(o.HostProfile) {
		toSerialize["host-profile"] = o.HostProfile
	}
	if !IsNil(o.HostProfileNumeric) {
		toSerialize["host-profile-numeric"] = o.HostProfileNumeric
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.InitiatorsUrl) {
		toSerialize["initiators-url"] = o.InitiatorsUrl
	}
	if !IsNil(o.Lun) {
		toSerialize["lun"] = o.Lun
	}
	if !IsNil(o.MappedId) {
		toSerialize["mapped-id"] = o.MappedId
	}
	if !IsNil(o.Nickname) {
		toSerialize["nickname"] = o.Nickname
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent-id"] = o.ParentId
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return toSerialize, nil
}

type NullableVolumeViewMappingsResourceInner struct {
	value *VolumeViewMappingsResourceInner
	isSet bool
}

func (v NullableVolumeViewMappingsResourceInner) Get() *VolumeViewMappingsResourceInner {
	return v.value
}

func (v *NullableVolumeViewMappingsResourceInner) Set(val *VolumeViewMappingsResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeViewMappingsResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeViewMappingsResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeViewMappingsResourceInner(val *VolumeViewMappingsResourceInner) *NullableVolumeViewMappingsResourceInner {
	return &NullableVolumeViewMappingsResourceInner{value: val, isSet: true}
}

func (v NullableVolumeViewMappingsResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeViewMappingsResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
