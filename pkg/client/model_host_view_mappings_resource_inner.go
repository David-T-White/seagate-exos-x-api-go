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

// checks if the HostViewMappingsResourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostViewMappingsResourceInner{}

// HostViewMappingsResourceInner struct for HostViewMappingsResourceInner
type HostViewMappingsResourceInner struct {
	ObjectName *string `json:"object-name,omitempty"`
	Meta       *string `json:"meta,omitempty"`
	// Access rights for this host
	Access *string `json:"access,omitempty"`
	// Access rights for this host( In numeric form )
	AccessNumeric *int64 `json:"access-numeric,omitempty"`
	// Logical Unit Number
	Lun   *string `json:"lun,omitempty"`
	Ports *string `json:"ports,omitempty"`
	// User-defined name for the volume
	Volume *string `json:"volume,omitempty"`
	// User-defined name for the volume
	VolumeName *string `json:"volume-name,omitempty"`
	// Unique serial number for the volume
	VolumeSerial *string                   `json:"volume-serial,omitempty"`
	VolumeView   []VolumeViewResourceInner `json:"volume-view,omitempty"`
}

// NewHostViewMappingsResourceInner instantiates a new HostViewMappingsResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostViewMappingsResourceInner() *HostViewMappingsResourceInner {
	this := HostViewMappingsResourceInner{}
	return &this
}

// NewHostViewMappingsResourceInnerWithDefaults instantiates a new HostViewMappingsResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostViewMappingsResourceInnerWithDefaults() *HostViewMappingsResourceInner {
	this := HostViewMappingsResourceInner{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *HostViewMappingsResourceInner) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostViewMappingsResourceInner) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *HostViewMappingsResourceInner) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *HostViewMappingsResourceInner) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *HostViewMappingsResourceInner) GetMeta() string {
	if o == nil || IsNil(o.Meta) {
		var ret string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostViewMappingsResourceInner) GetMetaOk() (*string, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *HostViewMappingsResourceInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given string and assigns it to the Meta field.
func (o *HostViewMappingsResourceInner) SetMeta(v string) {
	o.Meta = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *HostViewMappingsResourceInner) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostViewMappingsResourceInner) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *HostViewMappingsResourceInner) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *HostViewMappingsResourceInner) SetAccess(v string) {
	o.Access = &v
}

// GetAccessNumeric returns the AccessNumeric field value if set, zero value otherwise.
func (o *HostViewMappingsResourceInner) GetAccessNumeric() int64 {
	if o == nil || IsNil(o.AccessNumeric) {
		var ret int64
		return ret
	}
	return *o.AccessNumeric
}

// GetAccessNumericOk returns a tuple with the AccessNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostViewMappingsResourceInner) GetAccessNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessNumeric) {
		return nil, false
	}
	return o.AccessNumeric, true
}

// HasAccessNumeric returns a boolean if a field has been set.
func (o *HostViewMappingsResourceInner) HasAccessNumeric() bool {
	if o != nil && !IsNil(o.AccessNumeric) {
		return true
	}

	return false
}

// SetAccessNumeric gets a reference to the given int64 and assigns it to the AccessNumeric field.
func (o *HostViewMappingsResourceInner) SetAccessNumeric(v int64) {
	o.AccessNumeric = &v
}

// GetLun returns the Lun field value if set, zero value otherwise.
func (o *HostViewMappingsResourceInner) GetLun() string {
	if o == nil || IsNil(o.Lun) {
		var ret string
		return ret
	}
	return *o.Lun
}

// GetLunOk returns a tuple with the Lun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostViewMappingsResourceInner) GetLunOk() (*string, bool) {
	if o == nil || IsNil(o.Lun) {
		return nil, false
	}
	return o.Lun, true
}

// HasLun returns a boolean if a field has been set.
func (o *HostViewMappingsResourceInner) HasLun() bool {
	if o != nil && !IsNil(o.Lun) {
		return true
	}

	return false
}

// SetLun gets a reference to the given string and assigns it to the Lun field.
func (o *HostViewMappingsResourceInner) SetLun(v string) {
	o.Lun = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *HostViewMappingsResourceInner) GetPorts() string {
	if o == nil || IsNil(o.Ports) {
		var ret string
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostViewMappingsResourceInner) GetPortsOk() (*string, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *HostViewMappingsResourceInner) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given string and assigns it to the Ports field.
func (o *HostViewMappingsResourceInner) SetPorts(v string) {
	o.Ports = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *HostViewMappingsResourceInner) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostViewMappingsResourceInner) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *HostViewMappingsResourceInner) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *HostViewMappingsResourceInner) SetVolume(v string) {
	o.Volume = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *HostViewMappingsResourceInner) GetVolumeName() string {
	if o == nil || IsNil(o.VolumeName) {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostViewMappingsResourceInner) GetVolumeNameOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeName) {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *HostViewMappingsResourceInner) HasVolumeName() bool {
	if o != nil && !IsNil(o.VolumeName) {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *HostViewMappingsResourceInner) SetVolumeName(v string) {
	o.VolumeName = &v
}

// GetVolumeSerial returns the VolumeSerial field value if set, zero value otherwise.
func (o *HostViewMappingsResourceInner) GetVolumeSerial() string {
	if o == nil || IsNil(o.VolumeSerial) {
		var ret string
		return ret
	}
	return *o.VolumeSerial
}

// GetVolumeSerialOk returns a tuple with the VolumeSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostViewMappingsResourceInner) GetVolumeSerialOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeSerial) {
		return nil, false
	}
	return o.VolumeSerial, true
}

// HasVolumeSerial returns a boolean if a field has been set.
func (o *HostViewMappingsResourceInner) HasVolumeSerial() bool {
	if o != nil && !IsNil(o.VolumeSerial) {
		return true
	}

	return false
}

// SetVolumeSerial gets a reference to the given string and assigns it to the VolumeSerial field.
func (o *HostViewMappingsResourceInner) SetVolumeSerial(v string) {
	o.VolumeSerial = &v
}

// GetVolumeView returns the VolumeView field value if set, zero value otherwise.
func (o *HostViewMappingsResourceInner) GetVolumeView() []VolumeViewResourceInner {
	if o == nil || IsNil(o.VolumeView) {
		var ret []VolumeViewResourceInner
		return ret
	}
	return o.VolumeView
}

// GetVolumeViewOk returns a tuple with the VolumeView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostViewMappingsResourceInner) GetVolumeViewOk() ([]VolumeViewResourceInner, bool) {
	if o == nil || IsNil(o.VolumeView) {
		return nil, false
	}
	return o.VolumeView, true
}

// HasVolumeView returns a boolean if a field has been set.
func (o *HostViewMappingsResourceInner) HasVolumeView() bool {
	if o != nil && !IsNil(o.VolumeView) {
		return true
	}

	return false
}

// SetVolumeView gets a reference to the given []VolumeViewResourceInner and assigns it to the VolumeView field.
func (o *HostViewMappingsResourceInner) SetVolumeView(v []VolumeViewResourceInner) {
	o.VolumeView = v
}

func (o HostViewMappingsResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostViewMappingsResourceInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Lun) {
		toSerialize["lun"] = o.Lun
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !IsNil(o.VolumeName) {
		toSerialize["volume-name"] = o.VolumeName
	}
	if !IsNil(o.VolumeSerial) {
		toSerialize["volume-serial"] = o.VolumeSerial
	}
	if !IsNil(o.VolumeView) {
		toSerialize["volume-view"] = o.VolumeView
	}
	return toSerialize, nil
}

type NullableHostViewMappingsResourceInner struct {
	value *HostViewMappingsResourceInner
	isSet bool
}

func (v NullableHostViewMappingsResourceInner) Get() *HostViewMappingsResourceInner {
	return v.value
}

func (v *NullableHostViewMappingsResourceInner) Set(val *HostViewMappingsResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableHostViewMappingsResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableHostViewMappingsResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostViewMappingsResourceInner(val *HostViewMappingsResourceInner) *NullableHostViewMappingsResourceInner {
	return &NullableHostViewMappingsResourceInner{value: val, isSet: true}
}

func (v NullableHostViewMappingsResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostViewMappingsResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
