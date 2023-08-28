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

// HostsViewObject struct for HostsViewObject
type HostsViewObject struct {
	Status    []StatusResourceInner    `json:"status,omitempty"`
	HostsView []HostsViewResourceInner `json:"hosts-view,omitempty"`
}

// NewHostsViewObject instantiates a new HostsViewObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostsViewObject() *HostsViewObject {
	this := HostsViewObject{}
	return &this
}

// NewHostsViewObjectWithDefaults instantiates a new HostsViewObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostsViewObjectWithDefaults() *HostsViewObject {
	this := HostsViewObject{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HostsViewObject) GetStatus() []StatusResourceInner {
	if o == nil || isNil(o.Status) {
		var ret []StatusResourceInner
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostsViewObject) GetStatusOk() ([]StatusResourceInner, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HostsViewObject) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []StatusResourceInner and assigns it to the Status field.
func (o *HostsViewObject) SetStatus(v []StatusResourceInner) {
	o.Status = v
}

// GetHostsView returns the HostsView field value if set, zero value otherwise.
func (o *HostsViewObject) GetHostsView() []HostsViewResourceInner {
	if o == nil || isNil(o.HostsView) {
		var ret []HostsViewResourceInner
		return ret
	}
	return o.HostsView
}

// GetHostsViewOk returns a tuple with the HostsView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostsViewObject) GetHostsViewOk() ([]HostsViewResourceInner, bool) {
	if o == nil || isNil(o.HostsView) {
		return nil, false
	}
	return o.HostsView, true
}

// HasHostsView returns a boolean if a field has been set.
func (o *HostsViewObject) HasHostsView() bool {
	if o != nil && !isNil(o.HostsView) {
		return true
	}

	return false
}

// SetHostsView gets a reference to the given []HostsViewResourceInner and assigns it to the HostsView field.
func (o *HostsViewObject) SetHostsView(v []HostsViewResourceInner) {
	o.HostsView = v
}

func (o HostsViewObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.HostsView) {
		toSerialize["hosts-view"] = o.HostsView
	}
	return json.Marshal(toSerialize)
}

type NullableHostsViewObject struct {
	value *HostsViewObject
	isSet bool
}

func (v NullableHostsViewObject) Get() *HostsViewObject {
	return v.value
}

func (v *NullableHostsViewObject) Set(val *HostsViewObject) {
	v.value = val
	v.isSet = true
}

func (v NullableHostsViewObject) IsSet() bool {
	return v.isSet
}

func (v *NullableHostsViewObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostsViewObject(val *HostsViewObject) *NullableHostsViewObject {
	return &NullableHostsViewObject{value: val, isSet: true}
}

func (v NullableHostsViewObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostsViewObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
