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

// checks if the SystemObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemObject{}

// SystemObject struct for SystemObject
type SystemObject struct {
	Status []StatusResourceInner `json:"status,omitempty"`
	System []SystemResourceInner `json:"system,omitempty"`
}

// NewSystemObject instantiates a new SystemObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemObject() *SystemObject {
	this := SystemObject{}
	return &this
}

// NewSystemObjectWithDefaults instantiates a new SystemObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemObjectWithDefaults() *SystemObject {
	this := SystemObject{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SystemObject) GetStatus() []StatusResourceInner {
	if o == nil || IsNil(o.Status) {
		var ret []StatusResourceInner
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemObject) GetStatusOk() ([]StatusResourceInner, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SystemObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []StatusResourceInner and assigns it to the Status field.
func (o *SystemObject) SetStatus(v []StatusResourceInner) {
	o.Status = v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *SystemObject) GetSystem() []SystemResourceInner {
	if o == nil || IsNil(o.System) {
		var ret []SystemResourceInner
		return ret
	}
	return o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemObject) GetSystemOk() ([]SystemResourceInner, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *SystemObject) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given []SystemResourceInner and assigns it to the System field.
func (o *SystemObject) SetSystem(v []SystemResourceInner) {
	o.System = v
}

func (o SystemObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	return toSerialize, nil
}

type NullableSystemObject struct {
	value *SystemObject
	isSet bool
}

func (v NullableSystemObject) Get() *SystemObject {
	return v.value
}

func (v *NullableSystemObject) Set(val *SystemObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemObject(val *SystemObject) *NullableSystemObject {
	return &NullableSystemObject{value: val, isSet: true}
}

func (v NullableSystemObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
