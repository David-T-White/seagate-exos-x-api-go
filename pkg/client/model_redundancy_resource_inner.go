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

// checks if the RedundancyResourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedundancyResourceInner{}

// RedundancyResourceInner struct for RedundancyResourceInner
type RedundancyResourceInner struct {
	ObjectName               *string `json:"object-name,omitempty"`
	Meta                     *string `json:"meta,omitempty"`
	ControllerASerialNumber  *string `json:"controller-a-serial-number,omitempty"`
	ControllerAStatus        *string `json:"controller-a-status,omitempty"`
	ControllerAStatusNumeric *int64  `json:"controller-a-status-numeric,omitempty"`
	ControllerBSerialNumber  *string `json:"controller-b-serial-number,omitempty"`
	ControllerBStatus        *string `json:"controller-b-status,omitempty"`
	ControllerBStatusNumeric *int64  `json:"controller-b-status-numeric,omitempty"`
	LocalReady               *string `json:"local-ready,omitempty"`
	LocalReadyNumeric        *int64  `json:"local-ready-numeric,omitempty"`
	LocalReason              *string `json:"local-reason,omitempty"`
	// Identifies the availability of the partner MC
	OtherMCStatus *string `json:"other-MC-status,omitempty"`
	// Identifies the availability of the partner MC( In numeric form )
	OtherMCStatusNumeric *int64  `json:"other-MC-status-numeric,omitempty"`
	OtherReady           *string `json:"other-ready,omitempty"`
	OtherReadyNumeric    *int64  `json:"other-ready-numeric,omitempty"`
	OtherReason          *string `json:"other-reason,omitempty"`
	// Mode in which the controllers are operating
	RedundancyMode *string `json:"redundancy-mode,omitempty"`
	// Mode in which the controllers are operating( In numeric form )
	RedundancyModeNumeric *int64 `json:"redundancy-mode-numeric,omitempty"`
	// Current operational state of the controllers
	RedundancyStatus *string `json:"redundancy-status,omitempty"`
	// Current operational state of the controllers( In numeric form )
	RedundancyStatusNumeric *int64  `json:"redundancy-status-numeric,omitempty"`
	SystemReady             *string `json:"system-ready,omitempty"`
	SystemReadyNumeric      *int64  `json:"system-ready-numeric,omitempty"`
}

// NewRedundancyResourceInner instantiates a new RedundancyResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedundancyResourceInner() *RedundancyResourceInner {
	this := RedundancyResourceInner{}
	return &this
}

// NewRedundancyResourceInnerWithDefaults instantiates a new RedundancyResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedundancyResourceInnerWithDefaults() *RedundancyResourceInner {
	this := RedundancyResourceInner{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *RedundancyResourceInner) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetMeta() string {
	if o == nil || IsNil(o.Meta) {
		var ret string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetMetaOk() (*string, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given string and assigns it to the Meta field.
func (o *RedundancyResourceInner) SetMeta(v string) {
	o.Meta = &v
}

// GetControllerASerialNumber returns the ControllerASerialNumber field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetControllerASerialNumber() string {
	if o == nil || IsNil(o.ControllerASerialNumber) {
		var ret string
		return ret
	}
	return *o.ControllerASerialNumber
}

// GetControllerASerialNumberOk returns a tuple with the ControllerASerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetControllerASerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerASerialNumber) {
		return nil, false
	}
	return o.ControllerASerialNumber, true
}

// HasControllerASerialNumber returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasControllerASerialNumber() bool {
	if o != nil && !IsNil(o.ControllerASerialNumber) {
		return true
	}

	return false
}

// SetControllerASerialNumber gets a reference to the given string and assigns it to the ControllerASerialNumber field.
func (o *RedundancyResourceInner) SetControllerASerialNumber(v string) {
	o.ControllerASerialNumber = &v
}

// GetControllerAStatus returns the ControllerAStatus field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetControllerAStatus() string {
	if o == nil || IsNil(o.ControllerAStatus) {
		var ret string
		return ret
	}
	return *o.ControllerAStatus
}

// GetControllerAStatusOk returns a tuple with the ControllerAStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetControllerAStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerAStatus) {
		return nil, false
	}
	return o.ControllerAStatus, true
}

// HasControllerAStatus returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasControllerAStatus() bool {
	if o != nil && !IsNil(o.ControllerAStatus) {
		return true
	}

	return false
}

// SetControllerAStatus gets a reference to the given string and assigns it to the ControllerAStatus field.
func (o *RedundancyResourceInner) SetControllerAStatus(v string) {
	o.ControllerAStatus = &v
}

// GetControllerAStatusNumeric returns the ControllerAStatusNumeric field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetControllerAStatusNumeric() int64 {
	if o == nil || IsNil(o.ControllerAStatusNumeric) {
		var ret int64
		return ret
	}
	return *o.ControllerAStatusNumeric
}

// GetControllerAStatusNumericOk returns a tuple with the ControllerAStatusNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetControllerAStatusNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.ControllerAStatusNumeric) {
		return nil, false
	}
	return o.ControllerAStatusNumeric, true
}

// HasControllerAStatusNumeric returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasControllerAStatusNumeric() bool {
	if o != nil && !IsNil(o.ControllerAStatusNumeric) {
		return true
	}

	return false
}

// SetControllerAStatusNumeric gets a reference to the given int64 and assigns it to the ControllerAStatusNumeric field.
func (o *RedundancyResourceInner) SetControllerAStatusNumeric(v int64) {
	o.ControllerAStatusNumeric = &v
}

// GetControllerBSerialNumber returns the ControllerBSerialNumber field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetControllerBSerialNumber() string {
	if o == nil || IsNil(o.ControllerBSerialNumber) {
		var ret string
		return ret
	}
	return *o.ControllerBSerialNumber
}

// GetControllerBSerialNumberOk returns a tuple with the ControllerBSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetControllerBSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerBSerialNumber) {
		return nil, false
	}
	return o.ControllerBSerialNumber, true
}

// HasControllerBSerialNumber returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasControllerBSerialNumber() bool {
	if o != nil && !IsNil(o.ControllerBSerialNumber) {
		return true
	}

	return false
}

// SetControllerBSerialNumber gets a reference to the given string and assigns it to the ControllerBSerialNumber field.
func (o *RedundancyResourceInner) SetControllerBSerialNumber(v string) {
	o.ControllerBSerialNumber = &v
}

// GetControllerBStatus returns the ControllerBStatus field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetControllerBStatus() string {
	if o == nil || IsNil(o.ControllerBStatus) {
		var ret string
		return ret
	}
	return *o.ControllerBStatus
}

// GetControllerBStatusOk returns a tuple with the ControllerBStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetControllerBStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerBStatus) {
		return nil, false
	}
	return o.ControllerBStatus, true
}

// HasControllerBStatus returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasControllerBStatus() bool {
	if o != nil && !IsNil(o.ControllerBStatus) {
		return true
	}

	return false
}

// SetControllerBStatus gets a reference to the given string and assigns it to the ControllerBStatus field.
func (o *RedundancyResourceInner) SetControllerBStatus(v string) {
	o.ControllerBStatus = &v
}

// GetControllerBStatusNumeric returns the ControllerBStatusNumeric field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetControllerBStatusNumeric() int64 {
	if o == nil || IsNil(o.ControllerBStatusNumeric) {
		var ret int64
		return ret
	}
	return *o.ControllerBStatusNumeric
}

// GetControllerBStatusNumericOk returns a tuple with the ControllerBStatusNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetControllerBStatusNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.ControllerBStatusNumeric) {
		return nil, false
	}
	return o.ControllerBStatusNumeric, true
}

// HasControllerBStatusNumeric returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasControllerBStatusNumeric() bool {
	if o != nil && !IsNil(o.ControllerBStatusNumeric) {
		return true
	}

	return false
}

// SetControllerBStatusNumeric gets a reference to the given int64 and assigns it to the ControllerBStatusNumeric field.
func (o *RedundancyResourceInner) SetControllerBStatusNumeric(v int64) {
	o.ControllerBStatusNumeric = &v
}

// GetLocalReady returns the LocalReady field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetLocalReady() string {
	if o == nil || IsNil(o.LocalReady) {
		var ret string
		return ret
	}
	return *o.LocalReady
}

// GetLocalReadyOk returns a tuple with the LocalReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetLocalReadyOk() (*string, bool) {
	if o == nil || IsNil(o.LocalReady) {
		return nil, false
	}
	return o.LocalReady, true
}

// HasLocalReady returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasLocalReady() bool {
	if o != nil && !IsNil(o.LocalReady) {
		return true
	}

	return false
}

// SetLocalReady gets a reference to the given string and assigns it to the LocalReady field.
func (o *RedundancyResourceInner) SetLocalReady(v string) {
	o.LocalReady = &v
}

// GetLocalReadyNumeric returns the LocalReadyNumeric field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetLocalReadyNumeric() int64 {
	if o == nil || IsNil(o.LocalReadyNumeric) {
		var ret int64
		return ret
	}
	return *o.LocalReadyNumeric
}

// GetLocalReadyNumericOk returns a tuple with the LocalReadyNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetLocalReadyNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.LocalReadyNumeric) {
		return nil, false
	}
	return o.LocalReadyNumeric, true
}

// HasLocalReadyNumeric returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasLocalReadyNumeric() bool {
	if o != nil && !IsNil(o.LocalReadyNumeric) {
		return true
	}

	return false
}

// SetLocalReadyNumeric gets a reference to the given int64 and assigns it to the LocalReadyNumeric field.
func (o *RedundancyResourceInner) SetLocalReadyNumeric(v int64) {
	o.LocalReadyNumeric = &v
}

// GetLocalReason returns the LocalReason field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetLocalReason() string {
	if o == nil || IsNil(o.LocalReason) {
		var ret string
		return ret
	}
	return *o.LocalReason
}

// GetLocalReasonOk returns a tuple with the LocalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetLocalReasonOk() (*string, bool) {
	if o == nil || IsNil(o.LocalReason) {
		return nil, false
	}
	return o.LocalReason, true
}

// HasLocalReason returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasLocalReason() bool {
	if o != nil && !IsNil(o.LocalReason) {
		return true
	}

	return false
}

// SetLocalReason gets a reference to the given string and assigns it to the LocalReason field.
func (o *RedundancyResourceInner) SetLocalReason(v string) {
	o.LocalReason = &v
}

// GetOtherMCStatus returns the OtherMCStatus field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetOtherMCStatus() string {
	if o == nil || IsNil(o.OtherMCStatus) {
		var ret string
		return ret
	}
	return *o.OtherMCStatus
}

// GetOtherMCStatusOk returns a tuple with the OtherMCStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetOtherMCStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OtherMCStatus) {
		return nil, false
	}
	return o.OtherMCStatus, true
}

// HasOtherMCStatus returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasOtherMCStatus() bool {
	if o != nil && !IsNil(o.OtherMCStatus) {
		return true
	}

	return false
}

// SetOtherMCStatus gets a reference to the given string and assigns it to the OtherMCStatus field.
func (o *RedundancyResourceInner) SetOtherMCStatus(v string) {
	o.OtherMCStatus = &v
}

// GetOtherMCStatusNumeric returns the OtherMCStatusNumeric field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetOtherMCStatusNumeric() int64 {
	if o == nil || IsNil(o.OtherMCStatusNumeric) {
		var ret int64
		return ret
	}
	return *o.OtherMCStatusNumeric
}

// GetOtherMCStatusNumericOk returns a tuple with the OtherMCStatusNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetOtherMCStatusNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.OtherMCStatusNumeric) {
		return nil, false
	}
	return o.OtherMCStatusNumeric, true
}

// HasOtherMCStatusNumeric returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasOtherMCStatusNumeric() bool {
	if o != nil && !IsNil(o.OtherMCStatusNumeric) {
		return true
	}

	return false
}

// SetOtherMCStatusNumeric gets a reference to the given int64 and assigns it to the OtherMCStatusNumeric field.
func (o *RedundancyResourceInner) SetOtherMCStatusNumeric(v int64) {
	o.OtherMCStatusNumeric = &v
}

// GetOtherReady returns the OtherReady field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetOtherReady() string {
	if o == nil || IsNil(o.OtherReady) {
		var ret string
		return ret
	}
	return *o.OtherReady
}

// GetOtherReadyOk returns a tuple with the OtherReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetOtherReadyOk() (*string, bool) {
	if o == nil || IsNil(o.OtherReady) {
		return nil, false
	}
	return o.OtherReady, true
}

// HasOtherReady returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasOtherReady() bool {
	if o != nil && !IsNil(o.OtherReady) {
		return true
	}

	return false
}

// SetOtherReady gets a reference to the given string and assigns it to the OtherReady field.
func (o *RedundancyResourceInner) SetOtherReady(v string) {
	o.OtherReady = &v
}

// GetOtherReadyNumeric returns the OtherReadyNumeric field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetOtherReadyNumeric() int64 {
	if o == nil || IsNil(o.OtherReadyNumeric) {
		var ret int64
		return ret
	}
	return *o.OtherReadyNumeric
}

// GetOtherReadyNumericOk returns a tuple with the OtherReadyNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetOtherReadyNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.OtherReadyNumeric) {
		return nil, false
	}
	return o.OtherReadyNumeric, true
}

// HasOtherReadyNumeric returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasOtherReadyNumeric() bool {
	if o != nil && !IsNil(o.OtherReadyNumeric) {
		return true
	}

	return false
}

// SetOtherReadyNumeric gets a reference to the given int64 and assigns it to the OtherReadyNumeric field.
func (o *RedundancyResourceInner) SetOtherReadyNumeric(v int64) {
	o.OtherReadyNumeric = &v
}

// GetOtherReason returns the OtherReason field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetOtherReason() string {
	if o == nil || IsNil(o.OtherReason) {
		var ret string
		return ret
	}
	return *o.OtherReason
}

// GetOtherReasonOk returns a tuple with the OtherReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetOtherReasonOk() (*string, bool) {
	if o == nil || IsNil(o.OtherReason) {
		return nil, false
	}
	return o.OtherReason, true
}

// HasOtherReason returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasOtherReason() bool {
	if o != nil && !IsNil(o.OtherReason) {
		return true
	}

	return false
}

// SetOtherReason gets a reference to the given string and assigns it to the OtherReason field.
func (o *RedundancyResourceInner) SetOtherReason(v string) {
	o.OtherReason = &v
}

// GetRedundancyMode returns the RedundancyMode field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetRedundancyMode() string {
	if o == nil || IsNil(o.RedundancyMode) {
		var ret string
		return ret
	}
	return *o.RedundancyMode
}

// GetRedundancyModeOk returns a tuple with the RedundancyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetRedundancyModeOk() (*string, bool) {
	if o == nil || IsNil(o.RedundancyMode) {
		return nil, false
	}
	return o.RedundancyMode, true
}

// HasRedundancyMode returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasRedundancyMode() bool {
	if o != nil && !IsNil(o.RedundancyMode) {
		return true
	}

	return false
}

// SetRedundancyMode gets a reference to the given string and assigns it to the RedundancyMode field.
func (o *RedundancyResourceInner) SetRedundancyMode(v string) {
	o.RedundancyMode = &v
}

// GetRedundancyModeNumeric returns the RedundancyModeNumeric field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetRedundancyModeNumeric() int64 {
	if o == nil || IsNil(o.RedundancyModeNumeric) {
		var ret int64
		return ret
	}
	return *o.RedundancyModeNumeric
}

// GetRedundancyModeNumericOk returns a tuple with the RedundancyModeNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetRedundancyModeNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.RedundancyModeNumeric) {
		return nil, false
	}
	return o.RedundancyModeNumeric, true
}

// HasRedundancyModeNumeric returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasRedundancyModeNumeric() bool {
	if o != nil && !IsNil(o.RedundancyModeNumeric) {
		return true
	}

	return false
}

// SetRedundancyModeNumeric gets a reference to the given int64 and assigns it to the RedundancyModeNumeric field.
func (o *RedundancyResourceInner) SetRedundancyModeNumeric(v int64) {
	o.RedundancyModeNumeric = &v
}

// GetRedundancyStatus returns the RedundancyStatus field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetRedundancyStatus() string {
	if o == nil || IsNil(o.RedundancyStatus) {
		var ret string
		return ret
	}
	return *o.RedundancyStatus
}

// GetRedundancyStatusOk returns a tuple with the RedundancyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetRedundancyStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RedundancyStatus) {
		return nil, false
	}
	return o.RedundancyStatus, true
}

// HasRedundancyStatus returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasRedundancyStatus() bool {
	if o != nil && !IsNil(o.RedundancyStatus) {
		return true
	}

	return false
}

// SetRedundancyStatus gets a reference to the given string and assigns it to the RedundancyStatus field.
func (o *RedundancyResourceInner) SetRedundancyStatus(v string) {
	o.RedundancyStatus = &v
}

// GetRedundancyStatusNumeric returns the RedundancyStatusNumeric field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetRedundancyStatusNumeric() int64 {
	if o == nil || IsNil(o.RedundancyStatusNumeric) {
		var ret int64
		return ret
	}
	return *o.RedundancyStatusNumeric
}

// GetRedundancyStatusNumericOk returns a tuple with the RedundancyStatusNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetRedundancyStatusNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.RedundancyStatusNumeric) {
		return nil, false
	}
	return o.RedundancyStatusNumeric, true
}

// HasRedundancyStatusNumeric returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasRedundancyStatusNumeric() bool {
	if o != nil && !IsNil(o.RedundancyStatusNumeric) {
		return true
	}

	return false
}

// SetRedundancyStatusNumeric gets a reference to the given int64 and assigns it to the RedundancyStatusNumeric field.
func (o *RedundancyResourceInner) SetRedundancyStatusNumeric(v int64) {
	o.RedundancyStatusNumeric = &v
}

// GetSystemReady returns the SystemReady field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetSystemReady() string {
	if o == nil || IsNil(o.SystemReady) {
		var ret string
		return ret
	}
	return *o.SystemReady
}

// GetSystemReadyOk returns a tuple with the SystemReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetSystemReadyOk() (*string, bool) {
	if o == nil || IsNil(o.SystemReady) {
		return nil, false
	}
	return o.SystemReady, true
}

// HasSystemReady returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasSystemReady() bool {
	if o != nil && !IsNil(o.SystemReady) {
		return true
	}

	return false
}

// SetSystemReady gets a reference to the given string and assigns it to the SystemReady field.
func (o *RedundancyResourceInner) SetSystemReady(v string) {
	o.SystemReady = &v
}

// GetSystemReadyNumeric returns the SystemReadyNumeric field value if set, zero value otherwise.
func (o *RedundancyResourceInner) GetSystemReadyNumeric() int64 {
	if o == nil || IsNil(o.SystemReadyNumeric) {
		var ret int64
		return ret
	}
	return *o.SystemReadyNumeric
}

// GetSystemReadyNumericOk returns a tuple with the SystemReadyNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedundancyResourceInner) GetSystemReadyNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.SystemReadyNumeric) {
		return nil, false
	}
	return o.SystemReadyNumeric, true
}

// HasSystemReadyNumeric returns a boolean if a field has been set.
func (o *RedundancyResourceInner) HasSystemReadyNumeric() bool {
	if o != nil && !IsNil(o.SystemReadyNumeric) {
		return true
	}

	return false
}

// SetSystemReadyNumeric gets a reference to the given int64 and assigns it to the SystemReadyNumeric field.
func (o *RedundancyResourceInner) SetSystemReadyNumeric(v int64) {
	o.SystemReadyNumeric = &v
}

func (o RedundancyResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedundancyResourceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectName) {
		toSerialize["object-name"] = o.ObjectName
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.ControllerASerialNumber) {
		toSerialize["controller-a-serial-number"] = o.ControllerASerialNumber
	}
	if !IsNil(o.ControllerAStatus) {
		toSerialize["controller-a-status"] = o.ControllerAStatus
	}
	if !IsNil(o.ControllerAStatusNumeric) {
		toSerialize["controller-a-status-numeric"] = o.ControllerAStatusNumeric
	}
	if !IsNil(o.ControllerBSerialNumber) {
		toSerialize["controller-b-serial-number"] = o.ControllerBSerialNumber
	}
	if !IsNil(o.ControllerBStatus) {
		toSerialize["controller-b-status"] = o.ControllerBStatus
	}
	if !IsNil(o.ControllerBStatusNumeric) {
		toSerialize["controller-b-status-numeric"] = o.ControllerBStatusNumeric
	}
	if !IsNil(o.LocalReady) {
		toSerialize["local-ready"] = o.LocalReady
	}
	if !IsNil(o.LocalReadyNumeric) {
		toSerialize["local-ready-numeric"] = o.LocalReadyNumeric
	}
	if !IsNil(o.LocalReason) {
		toSerialize["local-reason"] = o.LocalReason
	}
	if !IsNil(o.OtherMCStatus) {
		toSerialize["other-MC-status"] = o.OtherMCStatus
	}
	if !IsNil(o.OtherMCStatusNumeric) {
		toSerialize["other-MC-status-numeric"] = o.OtherMCStatusNumeric
	}
	if !IsNil(o.OtherReady) {
		toSerialize["other-ready"] = o.OtherReady
	}
	if !IsNil(o.OtherReadyNumeric) {
		toSerialize["other-ready-numeric"] = o.OtherReadyNumeric
	}
	if !IsNil(o.OtherReason) {
		toSerialize["other-reason"] = o.OtherReason
	}
	if !IsNil(o.RedundancyMode) {
		toSerialize["redundancy-mode"] = o.RedundancyMode
	}
	if !IsNil(o.RedundancyModeNumeric) {
		toSerialize["redundancy-mode-numeric"] = o.RedundancyModeNumeric
	}
	if !IsNil(o.RedundancyStatus) {
		toSerialize["redundancy-status"] = o.RedundancyStatus
	}
	if !IsNil(o.RedundancyStatusNumeric) {
		toSerialize["redundancy-status-numeric"] = o.RedundancyStatusNumeric
	}
	if !IsNil(o.SystemReady) {
		toSerialize["system-ready"] = o.SystemReady
	}
	if !IsNil(o.SystemReadyNumeric) {
		toSerialize["system-ready-numeric"] = o.SystemReadyNumeric
	}
	return toSerialize, nil
}

type NullableRedundancyResourceInner struct {
	value *RedundancyResourceInner
	isSet bool
}

func (v NullableRedundancyResourceInner) Get() *RedundancyResourceInner {
	return v.value
}

func (v *NullableRedundancyResourceInner) Set(val *RedundancyResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRedundancyResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRedundancyResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedundancyResourceInner(val *RedundancyResourceInner) *NullableRedundancyResourceInner {
	return &NullableRedundancyResourceInner{value: val, isSet: true}
}

func (v NullableRedundancyResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedundancyResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
