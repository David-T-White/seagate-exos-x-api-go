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

// checks if the ControllersObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllersObject{}

// ControllersObject struct for ControllersObject
type ControllersObject struct {
	Status      []StatusResourceInner      `json:"status,omitempty"`
	Controllers []ControllersResourceInner `json:"controllers,omitempty"`
}

// NewControllersObject instantiates a new ControllersObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersObject() *ControllersObject {
	this := ControllersObject{}
	return &this
}

// NewControllersObjectWithDefaults instantiates a new ControllersObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersObjectWithDefaults() *ControllersObject {
	this := ControllersObject{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ControllersObject) GetStatus() []StatusResourceInner {
	if o == nil || IsNil(o.Status) {
		var ret []StatusResourceInner
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersObject) GetStatusOk() ([]StatusResourceInner, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ControllersObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []StatusResourceInner and assigns it to the Status field.
func (o *ControllersObject) SetStatus(v []StatusResourceInner) {
	o.Status = v
}

// GetControllers returns the Controllers field value if set, zero value otherwise.
func (o *ControllersObject) GetControllers() []ControllersResourceInner {
	if o == nil || IsNil(o.Controllers) {
		var ret []ControllersResourceInner
		return ret
	}
	return o.Controllers
}

// GetControllersOk returns a tuple with the Controllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersObject) GetControllersOk() ([]ControllersResourceInner, bool) {
	if o == nil || IsNil(o.Controllers) {
		return nil, false
	}
	return o.Controllers, true
}

// HasControllers returns a boolean if a field has been set.
func (o *ControllersObject) HasControllers() bool {
	if o != nil && !IsNil(o.Controllers) {
		return true
	}

	return false
}

// SetControllers gets a reference to the given []ControllersResourceInner and assigns it to the Controllers field.
func (o *ControllersObject) SetControllers(v []ControllersResourceInner) {
	o.Controllers = v
}

func (o ControllersObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllersObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Controllers) {
		toSerialize["controllers"] = o.Controllers
	}
	return toSerialize, nil
}

type NullableControllersObject struct {
	value *ControllersObject
	isSet bool
}

func (v NullableControllersObject) Get() *ControllersObject {
	return v.value
}

func (v *NullableControllersObject) Set(val *ControllersObject) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersObject) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersObject(val *ControllersObject) *NullableControllersObject {
	return &NullableControllersObject{value: val, isSet: true}
}

func (v NullableControllersObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
