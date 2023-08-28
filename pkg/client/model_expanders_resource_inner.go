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

// ExpandersResourceInner struct for ExpandersResourceInner
type ExpandersResourceInner struct {
	ObjectName  *string `json:"object-name,omitempty"`
	Meta        *string `json:"meta,omitempty"`
	DomId       *int64  `json:"dom-id,omitempty"`
	DrawerId    *int64  `json:"drawer-id,omitempty"`
	DurableId   *string `json:"durable-id,omitempty"`
	EnclosureId *int64  `json:"enclosure-id,omitempty"`
	// Extended status (bits)
	ExtendedStatus *string `json:"extended-status,omitempty"`
	// Firmware version of the FRU
	FwRevision           *string `json:"fw-revision,omitempty"`
	Health               *string `json:"health,omitempty"`
	HealthNumeric        *int64  `json:"health-numeric,omitempty"`
	HealthReason         *string `json:"health-reason,omitempty"`
	HealthRecommendation *string `json:"health-recommendation,omitempty"`
	Location             *string `json:"location,omitempty"`
	Name                 *string `json:"name,omitempty"`
	PathId               *string `json:"path-id,omitempty"`
	PathIdNumeric        *int64  `json:"path-id-numeric,omitempty"`
	Status               *string `json:"status,omitempty"`
	StatusNumeric        *int64  `json:"status-numeric,omitempty"`
}

// NewExpandersResourceInner instantiates a new ExpandersResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpandersResourceInner() *ExpandersResourceInner {
	this := ExpandersResourceInner{}
	return &this
}

// NewExpandersResourceInnerWithDefaults instantiates a new ExpandersResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpandersResourceInnerWithDefaults() *ExpandersResourceInner {
	this := ExpandersResourceInner{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetObjectName() string {
	if o == nil || isNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetObjectNameOk() (*string, bool) {
	if o == nil || isNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasObjectName() bool {
	if o != nil && !isNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ExpandersResourceInner) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetMeta() string {
	if o == nil || isNil(o.Meta) {
		var ret string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetMetaOk() (*string, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given string and assigns it to the Meta field.
func (o *ExpandersResourceInner) SetMeta(v string) {
	o.Meta = &v
}

// GetDomId returns the DomId field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetDomId() int64 {
	if o == nil || isNil(o.DomId) {
		var ret int64
		return ret
	}
	return *o.DomId
}

// GetDomIdOk returns a tuple with the DomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetDomIdOk() (*int64, bool) {
	if o == nil || isNil(o.DomId) {
		return nil, false
	}
	return o.DomId, true
}

// HasDomId returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasDomId() bool {
	if o != nil && !isNil(o.DomId) {
		return true
	}

	return false
}

// SetDomId gets a reference to the given int64 and assigns it to the DomId field.
func (o *ExpandersResourceInner) SetDomId(v int64) {
	o.DomId = &v
}

// GetDrawerId returns the DrawerId field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetDrawerId() int64 {
	if o == nil || isNil(o.DrawerId) {
		var ret int64
		return ret
	}
	return *o.DrawerId
}

// GetDrawerIdOk returns a tuple with the DrawerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetDrawerIdOk() (*int64, bool) {
	if o == nil || isNil(o.DrawerId) {
		return nil, false
	}
	return o.DrawerId, true
}

// HasDrawerId returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasDrawerId() bool {
	if o != nil && !isNil(o.DrawerId) {
		return true
	}

	return false
}

// SetDrawerId gets a reference to the given int64 and assigns it to the DrawerId field.
func (o *ExpandersResourceInner) SetDrawerId(v int64) {
	o.DrawerId = &v
}

// GetDurableId returns the DurableId field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetDurableId() string {
	if o == nil || isNil(o.DurableId) {
		var ret string
		return ret
	}
	return *o.DurableId
}

// GetDurableIdOk returns a tuple with the DurableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetDurableIdOk() (*string, bool) {
	if o == nil || isNil(o.DurableId) {
		return nil, false
	}
	return o.DurableId, true
}

// HasDurableId returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasDurableId() bool {
	if o != nil && !isNil(o.DurableId) {
		return true
	}

	return false
}

// SetDurableId gets a reference to the given string and assigns it to the DurableId field.
func (o *ExpandersResourceInner) SetDurableId(v string) {
	o.DurableId = &v
}

// GetEnclosureId returns the EnclosureId field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetEnclosureId() int64 {
	if o == nil || isNil(o.EnclosureId) {
		var ret int64
		return ret
	}
	return *o.EnclosureId
}

// GetEnclosureIdOk returns a tuple with the EnclosureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetEnclosureIdOk() (*int64, bool) {
	if o == nil || isNil(o.EnclosureId) {
		return nil, false
	}
	return o.EnclosureId, true
}

// HasEnclosureId returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasEnclosureId() bool {
	if o != nil && !isNil(o.EnclosureId) {
		return true
	}

	return false
}

// SetEnclosureId gets a reference to the given int64 and assigns it to the EnclosureId field.
func (o *ExpandersResourceInner) SetEnclosureId(v int64) {
	o.EnclosureId = &v
}

// GetExtendedStatus returns the ExtendedStatus field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetExtendedStatus() string {
	if o == nil || isNil(o.ExtendedStatus) {
		var ret string
		return ret
	}
	return *o.ExtendedStatus
}

// GetExtendedStatusOk returns a tuple with the ExtendedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetExtendedStatusOk() (*string, bool) {
	if o == nil || isNil(o.ExtendedStatus) {
		return nil, false
	}
	return o.ExtendedStatus, true
}

// HasExtendedStatus returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasExtendedStatus() bool {
	if o != nil && !isNil(o.ExtendedStatus) {
		return true
	}

	return false
}

// SetExtendedStatus gets a reference to the given string and assigns it to the ExtendedStatus field.
func (o *ExpandersResourceInner) SetExtendedStatus(v string) {
	o.ExtendedStatus = &v
}

// GetFwRevision returns the FwRevision field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetFwRevision() string {
	if o == nil || isNil(o.FwRevision) {
		var ret string
		return ret
	}
	return *o.FwRevision
}

// GetFwRevisionOk returns a tuple with the FwRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetFwRevisionOk() (*string, bool) {
	if o == nil || isNil(o.FwRevision) {
		return nil, false
	}
	return o.FwRevision, true
}

// HasFwRevision returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasFwRevision() bool {
	if o != nil && !isNil(o.FwRevision) {
		return true
	}

	return false
}

// SetFwRevision gets a reference to the given string and assigns it to the FwRevision field.
func (o *ExpandersResourceInner) SetFwRevision(v string) {
	o.FwRevision = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetHealth() string {
	if o == nil || isNil(o.Health) {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetHealthOk() (*string, bool) {
	if o == nil || isNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasHealth() bool {
	if o != nil && !isNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *ExpandersResourceInner) SetHealth(v string) {
	o.Health = &v
}

// GetHealthNumeric returns the HealthNumeric field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetHealthNumeric() int64 {
	if o == nil || isNil(o.HealthNumeric) {
		var ret int64
		return ret
	}
	return *o.HealthNumeric
}

// GetHealthNumericOk returns a tuple with the HealthNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetHealthNumericOk() (*int64, bool) {
	if o == nil || isNil(o.HealthNumeric) {
		return nil, false
	}
	return o.HealthNumeric, true
}

// HasHealthNumeric returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasHealthNumeric() bool {
	if o != nil && !isNil(o.HealthNumeric) {
		return true
	}

	return false
}

// SetHealthNumeric gets a reference to the given int64 and assigns it to the HealthNumeric field.
func (o *ExpandersResourceInner) SetHealthNumeric(v int64) {
	o.HealthNumeric = &v
}

// GetHealthReason returns the HealthReason field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetHealthReason() string {
	if o == nil || isNil(o.HealthReason) {
		var ret string
		return ret
	}
	return *o.HealthReason
}

// GetHealthReasonOk returns a tuple with the HealthReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetHealthReasonOk() (*string, bool) {
	if o == nil || isNil(o.HealthReason) {
		return nil, false
	}
	return o.HealthReason, true
}

// HasHealthReason returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasHealthReason() bool {
	if o != nil && !isNil(o.HealthReason) {
		return true
	}

	return false
}

// SetHealthReason gets a reference to the given string and assigns it to the HealthReason field.
func (o *ExpandersResourceInner) SetHealthReason(v string) {
	o.HealthReason = &v
}

// GetHealthRecommendation returns the HealthRecommendation field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetHealthRecommendation() string {
	if o == nil || isNil(o.HealthRecommendation) {
		var ret string
		return ret
	}
	return *o.HealthRecommendation
}

// GetHealthRecommendationOk returns a tuple with the HealthRecommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetHealthRecommendationOk() (*string, bool) {
	if o == nil || isNil(o.HealthRecommendation) {
		return nil, false
	}
	return o.HealthRecommendation, true
}

// HasHealthRecommendation returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasHealthRecommendation() bool {
	if o != nil && !isNil(o.HealthRecommendation) {
		return true
	}

	return false
}

// SetHealthRecommendation gets a reference to the given string and assigns it to the HealthRecommendation field.
func (o *ExpandersResourceInner) SetHealthRecommendation(v string) {
	o.HealthRecommendation = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetLocation() string {
	if o == nil || isNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetLocationOk() (*string, bool) {
	if o == nil || isNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ExpandersResourceInner) SetLocation(v string) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExpandersResourceInner) SetName(v string) {
	o.Name = &v
}

// GetPathId returns the PathId field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetPathId() string {
	if o == nil || isNil(o.PathId) {
		var ret string
		return ret
	}
	return *o.PathId
}

// GetPathIdOk returns a tuple with the PathId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetPathIdOk() (*string, bool) {
	if o == nil || isNil(o.PathId) {
		return nil, false
	}
	return o.PathId, true
}

// HasPathId returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasPathId() bool {
	if o != nil && !isNil(o.PathId) {
		return true
	}

	return false
}

// SetPathId gets a reference to the given string and assigns it to the PathId field.
func (o *ExpandersResourceInner) SetPathId(v string) {
	o.PathId = &v
}

// GetPathIdNumeric returns the PathIdNumeric field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetPathIdNumeric() int64 {
	if o == nil || isNil(o.PathIdNumeric) {
		var ret int64
		return ret
	}
	return *o.PathIdNumeric
}

// GetPathIdNumericOk returns a tuple with the PathIdNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetPathIdNumericOk() (*int64, bool) {
	if o == nil || isNil(o.PathIdNumeric) {
		return nil, false
	}
	return o.PathIdNumeric, true
}

// HasPathIdNumeric returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasPathIdNumeric() bool {
	if o != nil && !isNil(o.PathIdNumeric) {
		return true
	}

	return false
}

// SetPathIdNumeric gets a reference to the given int64 and assigns it to the PathIdNumeric field.
func (o *ExpandersResourceInner) SetPathIdNumeric(v int64) {
	o.PathIdNumeric = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExpandersResourceInner) SetStatus(v string) {
	o.Status = &v
}

// GetStatusNumeric returns the StatusNumeric field value if set, zero value otherwise.
func (o *ExpandersResourceInner) GetStatusNumeric() int64 {
	if o == nil || isNil(o.StatusNumeric) {
		var ret int64
		return ret
	}
	return *o.StatusNumeric
}

// GetStatusNumericOk returns a tuple with the StatusNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandersResourceInner) GetStatusNumericOk() (*int64, bool) {
	if o == nil || isNil(o.StatusNumeric) {
		return nil, false
	}
	return o.StatusNumeric, true
}

// HasStatusNumeric returns a boolean if a field has been set.
func (o *ExpandersResourceInner) HasStatusNumeric() bool {
	if o != nil && !isNil(o.StatusNumeric) {
		return true
	}

	return false
}

// SetStatusNumeric gets a reference to the given int64 and assigns it to the StatusNumeric field.
func (o *ExpandersResourceInner) SetStatusNumeric(v int64) {
	o.StatusNumeric = &v
}

func (o ExpandersResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ObjectName) {
		toSerialize["object-name"] = o.ObjectName
	}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.DomId) {
		toSerialize["dom-id"] = o.DomId
	}
	if !isNil(o.DrawerId) {
		toSerialize["drawer-id"] = o.DrawerId
	}
	if !isNil(o.DurableId) {
		toSerialize["durable-id"] = o.DurableId
	}
	if !isNil(o.EnclosureId) {
		toSerialize["enclosure-id"] = o.EnclosureId
	}
	if !isNil(o.ExtendedStatus) {
		toSerialize["extended-status"] = o.ExtendedStatus
	}
	if !isNil(o.FwRevision) {
		toSerialize["fw-revision"] = o.FwRevision
	}
	if !isNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !isNil(o.HealthNumeric) {
		toSerialize["health-numeric"] = o.HealthNumeric
	}
	if !isNil(o.HealthReason) {
		toSerialize["health-reason"] = o.HealthReason
	}
	if !isNil(o.HealthRecommendation) {
		toSerialize["health-recommendation"] = o.HealthRecommendation
	}
	if !isNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PathId) {
		toSerialize["path-id"] = o.PathId
	}
	if !isNil(o.PathIdNumeric) {
		toSerialize["path-id-numeric"] = o.PathIdNumeric
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.StatusNumeric) {
		toSerialize["status-numeric"] = o.StatusNumeric
	}
	return json.Marshal(toSerialize)
}

type NullableExpandersResourceInner struct {
	value *ExpandersResourceInner
	isSet bool
}

func (v NullableExpandersResourceInner) Get() *ExpandersResourceInner {
	return v.value
}

func (v *NullableExpandersResourceInner) Set(val *ExpandersResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandersResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandersResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandersResourceInner(val *ExpandersResourceInner) *NullableExpandersResourceInner {
	return &NullableExpandersResourceInner{value: val, isSet: true}
}

func (v NullableExpandersResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandersResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
