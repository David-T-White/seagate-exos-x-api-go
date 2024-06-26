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

// checks if the SasPortResourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SasPortResourceInner{}

// SasPortResourceInner struct for SasPortResourceInner
type SasPortResourceInner struct {
	ObjectName                *string `json:"object-name,omitempty"`
	Meta                      *string `json:"meta,omitempty"`
	ConfiguredTopology        *string `json:"configured-topology,omitempty"`
	ConfiguredTopologyNumeric *int64  `json:"configured-topology-numeric,omitempty"`
	// Number of currently active PHY lanes for the SAS port
	SasActiveLanes *int64 `json:"sas-active-lanes,omitempty"`
	// Number of disabled PHY lanes for the SAS port
	SasDisabledLanes *int64 `json:"sas-disabled-lanes,omitempty"`
	// Expected number of active PHY lanes for the SAS port
	SasLanesExpected *int64 `json:"sas-lanes-expected,omitempty"`
	// Number of PHY lanes for the SAS port
	Width *int64 `json:"width,omitempty"`
}

// NewSasPortResourceInner instantiates a new SasPortResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSasPortResourceInner() *SasPortResourceInner {
	this := SasPortResourceInner{}
	return &this
}

// NewSasPortResourceInnerWithDefaults instantiates a new SasPortResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSasPortResourceInnerWithDefaults() *SasPortResourceInner {
	this := SasPortResourceInner{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *SasPortResourceInner) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SasPortResourceInner) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *SasPortResourceInner) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *SasPortResourceInner) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SasPortResourceInner) GetMeta() string {
	if o == nil || IsNil(o.Meta) {
		var ret string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SasPortResourceInner) GetMetaOk() (*string, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SasPortResourceInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given string and assigns it to the Meta field.
func (o *SasPortResourceInner) SetMeta(v string) {
	o.Meta = &v
}

// GetConfiguredTopology returns the ConfiguredTopology field value if set, zero value otherwise.
func (o *SasPortResourceInner) GetConfiguredTopology() string {
	if o == nil || IsNil(o.ConfiguredTopology) {
		var ret string
		return ret
	}
	return *o.ConfiguredTopology
}

// GetConfiguredTopologyOk returns a tuple with the ConfiguredTopology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SasPortResourceInner) GetConfiguredTopologyOk() (*string, bool) {
	if o == nil || IsNil(o.ConfiguredTopology) {
		return nil, false
	}
	return o.ConfiguredTopology, true
}

// HasConfiguredTopology returns a boolean if a field has been set.
func (o *SasPortResourceInner) HasConfiguredTopology() bool {
	if o != nil && !IsNil(o.ConfiguredTopology) {
		return true
	}

	return false
}

// SetConfiguredTopology gets a reference to the given string and assigns it to the ConfiguredTopology field.
func (o *SasPortResourceInner) SetConfiguredTopology(v string) {
	o.ConfiguredTopology = &v
}

// GetConfiguredTopologyNumeric returns the ConfiguredTopologyNumeric field value if set, zero value otherwise.
func (o *SasPortResourceInner) GetConfiguredTopologyNumeric() int64 {
	if o == nil || IsNil(o.ConfiguredTopologyNumeric) {
		var ret int64
		return ret
	}
	return *o.ConfiguredTopologyNumeric
}

// GetConfiguredTopologyNumericOk returns a tuple with the ConfiguredTopologyNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SasPortResourceInner) GetConfiguredTopologyNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.ConfiguredTopologyNumeric) {
		return nil, false
	}
	return o.ConfiguredTopologyNumeric, true
}

// HasConfiguredTopologyNumeric returns a boolean if a field has been set.
func (o *SasPortResourceInner) HasConfiguredTopologyNumeric() bool {
	if o != nil && !IsNil(o.ConfiguredTopologyNumeric) {
		return true
	}

	return false
}

// SetConfiguredTopologyNumeric gets a reference to the given int64 and assigns it to the ConfiguredTopologyNumeric field.
func (o *SasPortResourceInner) SetConfiguredTopologyNumeric(v int64) {
	o.ConfiguredTopologyNumeric = &v
}

// GetSasActiveLanes returns the SasActiveLanes field value if set, zero value otherwise.
func (o *SasPortResourceInner) GetSasActiveLanes() int64 {
	if o == nil || IsNil(o.SasActiveLanes) {
		var ret int64
		return ret
	}
	return *o.SasActiveLanes
}

// GetSasActiveLanesOk returns a tuple with the SasActiveLanes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SasPortResourceInner) GetSasActiveLanesOk() (*int64, bool) {
	if o == nil || IsNil(o.SasActiveLanes) {
		return nil, false
	}
	return o.SasActiveLanes, true
}

// HasSasActiveLanes returns a boolean if a field has been set.
func (o *SasPortResourceInner) HasSasActiveLanes() bool {
	if o != nil && !IsNil(o.SasActiveLanes) {
		return true
	}

	return false
}

// SetSasActiveLanes gets a reference to the given int64 and assigns it to the SasActiveLanes field.
func (o *SasPortResourceInner) SetSasActiveLanes(v int64) {
	o.SasActiveLanes = &v
}

// GetSasDisabledLanes returns the SasDisabledLanes field value if set, zero value otherwise.
func (o *SasPortResourceInner) GetSasDisabledLanes() int64 {
	if o == nil || IsNil(o.SasDisabledLanes) {
		var ret int64
		return ret
	}
	return *o.SasDisabledLanes
}

// GetSasDisabledLanesOk returns a tuple with the SasDisabledLanes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SasPortResourceInner) GetSasDisabledLanesOk() (*int64, bool) {
	if o == nil || IsNil(o.SasDisabledLanes) {
		return nil, false
	}
	return o.SasDisabledLanes, true
}

// HasSasDisabledLanes returns a boolean if a field has been set.
func (o *SasPortResourceInner) HasSasDisabledLanes() bool {
	if o != nil && !IsNil(o.SasDisabledLanes) {
		return true
	}

	return false
}

// SetSasDisabledLanes gets a reference to the given int64 and assigns it to the SasDisabledLanes field.
func (o *SasPortResourceInner) SetSasDisabledLanes(v int64) {
	o.SasDisabledLanes = &v
}

// GetSasLanesExpected returns the SasLanesExpected field value if set, zero value otherwise.
func (o *SasPortResourceInner) GetSasLanesExpected() int64 {
	if o == nil || IsNil(o.SasLanesExpected) {
		var ret int64
		return ret
	}
	return *o.SasLanesExpected
}

// GetSasLanesExpectedOk returns a tuple with the SasLanesExpected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SasPortResourceInner) GetSasLanesExpectedOk() (*int64, bool) {
	if o == nil || IsNil(o.SasLanesExpected) {
		return nil, false
	}
	return o.SasLanesExpected, true
}

// HasSasLanesExpected returns a boolean if a field has been set.
func (o *SasPortResourceInner) HasSasLanesExpected() bool {
	if o != nil && !IsNil(o.SasLanesExpected) {
		return true
	}

	return false
}

// SetSasLanesExpected gets a reference to the given int64 and assigns it to the SasLanesExpected field.
func (o *SasPortResourceInner) SetSasLanesExpected(v int64) {
	o.SasLanesExpected = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *SasPortResourceInner) GetWidth() int64 {
	if o == nil || IsNil(o.Width) {
		var ret int64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SasPortResourceInner) GetWidthOk() (*int64, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *SasPortResourceInner) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int64 and assigns it to the Width field.
func (o *SasPortResourceInner) SetWidth(v int64) {
	o.Width = &v
}

func (o SasPortResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SasPortResourceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectName) {
		toSerialize["object-name"] = o.ObjectName
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.ConfiguredTopology) {
		toSerialize["configured-topology"] = o.ConfiguredTopology
	}
	if !IsNil(o.ConfiguredTopologyNumeric) {
		toSerialize["configured-topology-numeric"] = o.ConfiguredTopologyNumeric
	}
	if !IsNil(o.SasActiveLanes) {
		toSerialize["sas-active-lanes"] = o.SasActiveLanes
	}
	if !IsNil(o.SasDisabledLanes) {
		toSerialize["sas-disabled-lanes"] = o.SasDisabledLanes
	}
	if !IsNil(o.SasLanesExpected) {
		toSerialize["sas-lanes-expected"] = o.SasLanesExpected
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullableSasPortResourceInner struct {
	value *SasPortResourceInner
	isSet bool
}

func (v NullableSasPortResourceInner) Get() *SasPortResourceInner {
	return v.value
}

func (v *NullableSasPortResourceInner) Set(val *SasPortResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSasPortResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSasPortResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSasPortResourceInner(val *SasPortResourceInner) *NullableSasPortResourceInner {
	return &NullableSasPortResourceInner{value: val, isSet: true}
}

func (v NullableSasPortResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSasPortResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
