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

// checks if the SystemResourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemResourceInner{}

// SystemResourceInner struct for SystemResourceInner
type SystemResourceInner struct {
	ObjectName               *string `json:"object-name,omitempty"`
	Meta                     *string `json:"meta,omitempty"`
	ConfigurationSelector    *string `json:"configuration-selector,omitempty"`
	CurrentNodeWwn           *string `json:"current-node-wwn,omitempty"`
	EnclosureCount           *int64  `json:"enclosure-count,omitempty"`
	FdeSecurityStatus        *string `json:"fde-security-status,omitempty"`
	FdeSecurityStatusNumeric *int64  `json:"fde-security-status-numeric,omitempty"`
	Health                   *string `json:"health,omitempty"`
	HealthNumeric            *int64  `json:"health-numeric,omitempty"`
	HealthReason             *string `json:"health-reason,omitempty"`
	// Serial number of the enclosure
	MidplaneSerialNumber *string `json:"midplane-serial-number,omitempty"`
	// Identifies the availability of the partner MC
	OtherMCStatus *string `json:"other-MC-status,omitempty"`
	// Identifies the availability of the partner MC( In numeric form )
	OtherMCStatusNumeric *int64  `json:"other-MC-status-numeric,omitempty"`
	PfuStatus            *string `json:"pfuStatus,omitempty"`
	PfuStatusNumeric     *int64  `json:"pfuStatus-numeric,omitempty"`
	// HW Platform Brand
	PlatformBrand *string `json:"platform-brand,omitempty"`
	// HW Platform Brand( In numeric form )
	PlatformBrandNumeric *int64 `json:"platform-brand-numeric,omitempty"`
	// HW Platform Type
	PlatformType *string `json:"platform-type,omitempty"`
	// HW Platform Type( In numeric form )
	PlatformTypeNumeric *int64  `json:"platform-type-numeric,omitempty"`
	ProductBrand        *string `json:"product-brand,omitempty"`
	ProductId           *string `json:"product-id,omitempty"`
	// SCSI Product ID
	ScsiProductId *string `json:"scsi-product-id,omitempty"`
	// SCSI vendor name
	ScsiVendorId                    *string `json:"scsi-vendor-id,omitempty"`
	SecuritySystemManagement        *string `json:"security-system-management,omitempty"`
	SecuritySystemManagementNumeric *int64  `json:"security-system-management-numeric,omitempty"`
	SupportedLocales                *string `json:"supported-locales,omitempty"`
	// User-defined contact for this system
	SystemContact     *string `json:"system-contact,omitempty"`
	SystemInformation *string `json:"system-information,omitempty"`
	// User-defined location of the system
	SystemLocation *string `json:"system-location,omitempty"`
	// User-defined name for the system
	SystemName *string `json:"system-name,omitempty"`
	// The resource URL
	Url        *string                   `json:"url,omitempty"`
	VendorName *string                   `json:"vendor-name,omitempty"`
	Redundancy []RedundancyResourceInner `json:"redundancy,omitempty"`
}

// NewSystemResourceInner instantiates a new SystemResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemResourceInner() *SystemResourceInner {
	this := SystemResourceInner{}
	return &this
}

// NewSystemResourceInnerWithDefaults instantiates a new SystemResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemResourceInnerWithDefaults() *SystemResourceInner {
	this := SystemResourceInner{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *SystemResourceInner) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *SystemResourceInner) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *SystemResourceInner) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SystemResourceInner) GetMeta() string {
	if o == nil || IsNil(o.Meta) {
		var ret string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetMetaOk() (*string, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SystemResourceInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given string and assigns it to the Meta field.
func (o *SystemResourceInner) SetMeta(v string) {
	o.Meta = &v
}

// GetConfigurationSelector returns the ConfigurationSelector field value if set, zero value otherwise.
func (o *SystemResourceInner) GetConfigurationSelector() string {
	if o == nil || IsNil(o.ConfigurationSelector) {
		var ret string
		return ret
	}
	return *o.ConfigurationSelector
}

// GetConfigurationSelectorOk returns a tuple with the ConfigurationSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetConfigurationSelectorOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationSelector) {
		return nil, false
	}
	return o.ConfigurationSelector, true
}

// HasConfigurationSelector returns a boolean if a field has been set.
func (o *SystemResourceInner) HasConfigurationSelector() bool {
	if o != nil && !IsNil(o.ConfigurationSelector) {
		return true
	}

	return false
}

// SetConfigurationSelector gets a reference to the given string and assigns it to the ConfigurationSelector field.
func (o *SystemResourceInner) SetConfigurationSelector(v string) {
	o.ConfigurationSelector = &v
}

// GetCurrentNodeWwn returns the CurrentNodeWwn field value if set, zero value otherwise.
func (o *SystemResourceInner) GetCurrentNodeWwn() string {
	if o == nil || IsNil(o.CurrentNodeWwn) {
		var ret string
		return ret
	}
	return *o.CurrentNodeWwn
}

// GetCurrentNodeWwnOk returns a tuple with the CurrentNodeWwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetCurrentNodeWwnOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentNodeWwn) {
		return nil, false
	}
	return o.CurrentNodeWwn, true
}

// HasCurrentNodeWwn returns a boolean if a field has been set.
func (o *SystemResourceInner) HasCurrentNodeWwn() bool {
	if o != nil && !IsNil(o.CurrentNodeWwn) {
		return true
	}

	return false
}

// SetCurrentNodeWwn gets a reference to the given string and assigns it to the CurrentNodeWwn field.
func (o *SystemResourceInner) SetCurrentNodeWwn(v string) {
	o.CurrentNodeWwn = &v
}

// GetEnclosureCount returns the EnclosureCount field value if set, zero value otherwise.
func (o *SystemResourceInner) GetEnclosureCount() int64 {
	if o == nil || IsNil(o.EnclosureCount) {
		var ret int64
		return ret
	}
	return *o.EnclosureCount
}

// GetEnclosureCountOk returns a tuple with the EnclosureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetEnclosureCountOk() (*int64, bool) {
	if o == nil || IsNil(o.EnclosureCount) {
		return nil, false
	}
	return o.EnclosureCount, true
}

// HasEnclosureCount returns a boolean if a field has been set.
func (o *SystemResourceInner) HasEnclosureCount() bool {
	if o != nil && !IsNil(o.EnclosureCount) {
		return true
	}

	return false
}

// SetEnclosureCount gets a reference to the given int64 and assigns it to the EnclosureCount field.
func (o *SystemResourceInner) SetEnclosureCount(v int64) {
	o.EnclosureCount = &v
}

// GetFdeSecurityStatus returns the FdeSecurityStatus field value if set, zero value otherwise.
func (o *SystemResourceInner) GetFdeSecurityStatus() string {
	if o == nil || IsNil(o.FdeSecurityStatus) {
		var ret string
		return ret
	}
	return *o.FdeSecurityStatus
}

// GetFdeSecurityStatusOk returns a tuple with the FdeSecurityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetFdeSecurityStatusOk() (*string, bool) {
	if o == nil || IsNil(o.FdeSecurityStatus) {
		return nil, false
	}
	return o.FdeSecurityStatus, true
}

// HasFdeSecurityStatus returns a boolean if a field has been set.
func (o *SystemResourceInner) HasFdeSecurityStatus() bool {
	if o != nil && !IsNil(o.FdeSecurityStatus) {
		return true
	}

	return false
}

// SetFdeSecurityStatus gets a reference to the given string and assigns it to the FdeSecurityStatus field.
func (o *SystemResourceInner) SetFdeSecurityStatus(v string) {
	o.FdeSecurityStatus = &v
}

// GetFdeSecurityStatusNumeric returns the FdeSecurityStatusNumeric field value if set, zero value otherwise.
func (o *SystemResourceInner) GetFdeSecurityStatusNumeric() int64 {
	if o == nil || IsNil(o.FdeSecurityStatusNumeric) {
		var ret int64
		return ret
	}
	return *o.FdeSecurityStatusNumeric
}

// GetFdeSecurityStatusNumericOk returns a tuple with the FdeSecurityStatusNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetFdeSecurityStatusNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.FdeSecurityStatusNumeric) {
		return nil, false
	}
	return o.FdeSecurityStatusNumeric, true
}

// HasFdeSecurityStatusNumeric returns a boolean if a field has been set.
func (o *SystemResourceInner) HasFdeSecurityStatusNumeric() bool {
	if o != nil && !IsNil(o.FdeSecurityStatusNumeric) {
		return true
	}

	return false
}

// SetFdeSecurityStatusNumeric gets a reference to the given int64 and assigns it to the FdeSecurityStatusNumeric field.
func (o *SystemResourceInner) SetFdeSecurityStatusNumeric(v int64) {
	o.FdeSecurityStatusNumeric = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *SystemResourceInner) GetHealth() string {
	if o == nil || IsNil(o.Health) {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetHealthOk() (*string, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *SystemResourceInner) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *SystemResourceInner) SetHealth(v string) {
	o.Health = &v
}

// GetHealthNumeric returns the HealthNumeric field value if set, zero value otherwise.
func (o *SystemResourceInner) GetHealthNumeric() int64 {
	if o == nil || IsNil(o.HealthNumeric) {
		var ret int64
		return ret
	}
	return *o.HealthNumeric
}

// GetHealthNumericOk returns a tuple with the HealthNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetHealthNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.HealthNumeric) {
		return nil, false
	}
	return o.HealthNumeric, true
}

// HasHealthNumeric returns a boolean if a field has been set.
func (o *SystemResourceInner) HasHealthNumeric() bool {
	if o != nil && !IsNil(o.HealthNumeric) {
		return true
	}

	return false
}

// SetHealthNumeric gets a reference to the given int64 and assigns it to the HealthNumeric field.
func (o *SystemResourceInner) SetHealthNumeric(v int64) {
	o.HealthNumeric = &v
}

// GetHealthReason returns the HealthReason field value if set, zero value otherwise.
func (o *SystemResourceInner) GetHealthReason() string {
	if o == nil || IsNil(o.HealthReason) {
		var ret string
		return ret
	}
	return *o.HealthReason
}

// GetHealthReasonOk returns a tuple with the HealthReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetHealthReasonOk() (*string, bool) {
	if o == nil || IsNil(o.HealthReason) {
		return nil, false
	}
	return o.HealthReason, true
}

// HasHealthReason returns a boolean if a field has been set.
func (o *SystemResourceInner) HasHealthReason() bool {
	if o != nil && !IsNil(o.HealthReason) {
		return true
	}

	return false
}

// SetHealthReason gets a reference to the given string and assigns it to the HealthReason field.
func (o *SystemResourceInner) SetHealthReason(v string) {
	o.HealthReason = &v
}

// GetMidplaneSerialNumber returns the MidplaneSerialNumber field value if set, zero value otherwise.
func (o *SystemResourceInner) GetMidplaneSerialNumber() string {
	if o == nil || IsNil(o.MidplaneSerialNumber) {
		var ret string
		return ret
	}
	return *o.MidplaneSerialNumber
}

// GetMidplaneSerialNumberOk returns a tuple with the MidplaneSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetMidplaneSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.MidplaneSerialNumber) {
		return nil, false
	}
	return o.MidplaneSerialNumber, true
}

// HasMidplaneSerialNumber returns a boolean if a field has been set.
func (o *SystemResourceInner) HasMidplaneSerialNumber() bool {
	if o != nil && !IsNil(o.MidplaneSerialNumber) {
		return true
	}

	return false
}

// SetMidplaneSerialNumber gets a reference to the given string and assigns it to the MidplaneSerialNumber field.
func (o *SystemResourceInner) SetMidplaneSerialNumber(v string) {
	o.MidplaneSerialNumber = &v
}

// GetOtherMCStatus returns the OtherMCStatus field value if set, zero value otherwise.
func (o *SystemResourceInner) GetOtherMCStatus() string {
	if o == nil || IsNil(o.OtherMCStatus) {
		var ret string
		return ret
	}
	return *o.OtherMCStatus
}

// GetOtherMCStatusOk returns a tuple with the OtherMCStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetOtherMCStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OtherMCStatus) {
		return nil, false
	}
	return o.OtherMCStatus, true
}

// HasOtherMCStatus returns a boolean if a field has been set.
func (o *SystemResourceInner) HasOtherMCStatus() bool {
	if o != nil && !IsNil(o.OtherMCStatus) {
		return true
	}

	return false
}

// SetOtherMCStatus gets a reference to the given string and assigns it to the OtherMCStatus field.
func (o *SystemResourceInner) SetOtherMCStatus(v string) {
	o.OtherMCStatus = &v
}

// GetOtherMCStatusNumeric returns the OtherMCStatusNumeric field value if set, zero value otherwise.
func (o *SystemResourceInner) GetOtherMCStatusNumeric() int64 {
	if o == nil || IsNil(o.OtherMCStatusNumeric) {
		var ret int64
		return ret
	}
	return *o.OtherMCStatusNumeric
}

// GetOtherMCStatusNumericOk returns a tuple with the OtherMCStatusNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetOtherMCStatusNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.OtherMCStatusNumeric) {
		return nil, false
	}
	return o.OtherMCStatusNumeric, true
}

// HasOtherMCStatusNumeric returns a boolean if a field has been set.
func (o *SystemResourceInner) HasOtherMCStatusNumeric() bool {
	if o != nil && !IsNil(o.OtherMCStatusNumeric) {
		return true
	}

	return false
}

// SetOtherMCStatusNumeric gets a reference to the given int64 and assigns it to the OtherMCStatusNumeric field.
func (o *SystemResourceInner) SetOtherMCStatusNumeric(v int64) {
	o.OtherMCStatusNumeric = &v
}

// GetPfuStatus returns the PfuStatus field value if set, zero value otherwise.
func (o *SystemResourceInner) GetPfuStatus() string {
	if o == nil || IsNil(o.PfuStatus) {
		var ret string
		return ret
	}
	return *o.PfuStatus
}

// GetPfuStatusOk returns a tuple with the PfuStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetPfuStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PfuStatus) {
		return nil, false
	}
	return o.PfuStatus, true
}

// HasPfuStatus returns a boolean if a field has been set.
func (o *SystemResourceInner) HasPfuStatus() bool {
	if o != nil && !IsNil(o.PfuStatus) {
		return true
	}

	return false
}

// SetPfuStatus gets a reference to the given string and assigns it to the PfuStatus field.
func (o *SystemResourceInner) SetPfuStatus(v string) {
	o.PfuStatus = &v
}

// GetPfuStatusNumeric returns the PfuStatusNumeric field value if set, zero value otherwise.
func (o *SystemResourceInner) GetPfuStatusNumeric() int64 {
	if o == nil || IsNil(o.PfuStatusNumeric) {
		var ret int64
		return ret
	}
	return *o.PfuStatusNumeric
}

// GetPfuStatusNumericOk returns a tuple with the PfuStatusNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetPfuStatusNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.PfuStatusNumeric) {
		return nil, false
	}
	return o.PfuStatusNumeric, true
}

// HasPfuStatusNumeric returns a boolean if a field has been set.
func (o *SystemResourceInner) HasPfuStatusNumeric() bool {
	if o != nil && !IsNil(o.PfuStatusNumeric) {
		return true
	}

	return false
}

// SetPfuStatusNumeric gets a reference to the given int64 and assigns it to the PfuStatusNumeric field.
func (o *SystemResourceInner) SetPfuStatusNumeric(v int64) {
	o.PfuStatusNumeric = &v
}

// GetPlatformBrand returns the PlatformBrand field value if set, zero value otherwise.
func (o *SystemResourceInner) GetPlatformBrand() string {
	if o == nil || IsNil(o.PlatformBrand) {
		var ret string
		return ret
	}
	return *o.PlatformBrand
}

// GetPlatformBrandOk returns a tuple with the PlatformBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetPlatformBrandOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformBrand) {
		return nil, false
	}
	return o.PlatformBrand, true
}

// HasPlatformBrand returns a boolean if a field has been set.
func (o *SystemResourceInner) HasPlatformBrand() bool {
	if o != nil && !IsNil(o.PlatformBrand) {
		return true
	}

	return false
}

// SetPlatformBrand gets a reference to the given string and assigns it to the PlatformBrand field.
func (o *SystemResourceInner) SetPlatformBrand(v string) {
	o.PlatformBrand = &v
}

// GetPlatformBrandNumeric returns the PlatformBrandNumeric field value if set, zero value otherwise.
func (o *SystemResourceInner) GetPlatformBrandNumeric() int64 {
	if o == nil || IsNil(o.PlatformBrandNumeric) {
		var ret int64
		return ret
	}
	return *o.PlatformBrandNumeric
}

// GetPlatformBrandNumericOk returns a tuple with the PlatformBrandNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetPlatformBrandNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.PlatformBrandNumeric) {
		return nil, false
	}
	return o.PlatformBrandNumeric, true
}

// HasPlatformBrandNumeric returns a boolean if a field has been set.
func (o *SystemResourceInner) HasPlatformBrandNumeric() bool {
	if o != nil && !IsNil(o.PlatformBrandNumeric) {
		return true
	}

	return false
}

// SetPlatformBrandNumeric gets a reference to the given int64 and assigns it to the PlatformBrandNumeric field.
func (o *SystemResourceInner) SetPlatformBrandNumeric(v int64) {
	o.PlatformBrandNumeric = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *SystemResourceInner) GetPlatformType() string {
	if o == nil || IsNil(o.PlatformType) {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetPlatformTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformType) {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *SystemResourceInner) HasPlatformType() bool {
	if o != nil && !IsNil(o.PlatformType) {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *SystemResourceInner) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetPlatformTypeNumeric returns the PlatformTypeNumeric field value if set, zero value otherwise.
func (o *SystemResourceInner) GetPlatformTypeNumeric() int64 {
	if o == nil || IsNil(o.PlatformTypeNumeric) {
		var ret int64
		return ret
	}
	return *o.PlatformTypeNumeric
}

// GetPlatformTypeNumericOk returns a tuple with the PlatformTypeNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetPlatformTypeNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.PlatformTypeNumeric) {
		return nil, false
	}
	return o.PlatformTypeNumeric, true
}

// HasPlatformTypeNumeric returns a boolean if a field has been set.
func (o *SystemResourceInner) HasPlatformTypeNumeric() bool {
	if o != nil && !IsNil(o.PlatformTypeNumeric) {
		return true
	}

	return false
}

// SetPlatformTypeNumeric gets a reference to the given int64 and assigns it to the PlatformTypeNumeric field.
func (o *SystemResourceInner) SetPlatformTypeNumeric(v int64) {
	o.PlatformTypeNumeric = &v
}

// GetProductBrand returns the ProductBrand field value if set, zero value otherwise.
func (o *SystemResourceInner) GetProductBrand() string {
	if o == nil || IsNil(o.ProductBrand) {
		var ret string
		return ret
	}
	return *o.ProductBrand
}

// GetProductBrandOk returns a tuple with the ProductBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetProductBrandOk() (*string, bool) {
	if o == nil || IsNil(o.ProductBrand) {
		return nil, false
	}
	return o.ProductBrand, true
}

// HasProductBrand returns a boolean if a field has been set.
func (o *SystemResourceInner) HasProductBrand() bool {
	if o != nil && !IsNil(o.ProductBrand) {
		return true
	}

	return false
}

// SetProductBrand gets a reference to the given string and assigns it to the ProductBrand field.
func (o *SystemResourceInner) SetProductBrand(v string) {
	o.ProductBrand = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *SystemResourceInner) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *SystemResourceInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *SystemResourceInner) SetProductId(v string) {
	o.ProductId = &v
}

// GetScsiProductId returns the ScsiProductId field value if set, zero value otherwise.
func (o *SystemResourceInner) GetScsiProductId() string {
	if o == nil || IsNil(o.ScsiProductId) {
		var ret string
		return ret
	}
	return *o.ScsiProductId
}

// GetScsiProductIdOk returns a tuple with the ScsiProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetScsiProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScsiProductId) {
		return nil, false
	}
	return o.ScsiProductId, true
}

// HasScsiProductId returns a boolean if a field has been set.
func (o *SystemResourceInner) HasScsiProductId() bool {
	if o != nil && !IsNil(o.ScsiProductId) {
		return true
	}

	return false
}

// SetScsiProductId gets a reference to the given string and assigns it to the ScsiProductId field.
func (o *SystemResourceInner) SetScsiProductId(v string) {
	o.ScsiProductId = &v
}

// GetScsiVendorId returns the ScsiVendorId field value if set, zero value otherwise.
func (o *SystemResourceInner) GetScsiVendorId() string {
	if o == nil || IsNil(o.ScsiVendorId) {
		var ret string
		return ret
	}
	return *o.ScsiVendorId
}

// GetScsiVendorIdOk returns a tuple with the ScsiVendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetScsiVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScsiVendorId) {
		return nil, false
	}
	return o.ScsiVendorId, true
}

// HasScsiVendorId returns a boolean if a field has been set.
func (o *SystemResourceInner) HasScsiVendorId() bool {
	if o != nil && !IsNil(o.ScsiVendorId) {
		return true
	}

	return false
}

// SetScsiVendorId gets a reference to the given string and assigns it to the ScsiVendorId field.
func (o *SystemResourceInner) SetScsiVendorId(v string) {
	o.ScsiVendorId = &v
}

// GetSecuritySystemManagement returns the SecuritySystemManagement field value if set, zero value otherwise.
func (o *SystemResourceInner) GetSecuritySystemManagement() string {
	if o == nil || IsNil(o.SecuritySystemManagement) {
		var ret string
		return ret
	}
	return *o.SecuritySystemManagement
}

// GetSecuritySystemManagementOk returns a tuple with the SecuritySystemManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetSecuritySystemManagementOk() (*string, bool) {
	if o == nil || IsNil(o.SecuritySystemManagement) {
		return nil, false
	}
	return o.SecuritySystemManagement, true
}

// HasSecuritySystemManagement returns a boolean if a field has been set.
func (o *SystemResourceInner) HasSecuritySystemManagement() bool {
	if o != nil && !IsNil(o.SecuritySystemManagement) {
		return true
	}

	return false
}

// SetSecuritySystemManagement gets a reference to the given string and assigns it to the SecuritySystemManagement field.
func (o *SystemResourceInner) SetSecuritySystemManagement(v string) {
	o.SecuritySystemManagement = &v
}

// GetSecuritySystemManagementNumeric returns the SecuritySystemManagementNumeric field value if set, zero value otherwise.
func (o *SystemResourceInner) GetSecuritySystemManagementNumeric() int64 {
	if o == nil || IsNil(o.SecuritySystemManagementNumeric) {
		var ret int64
		return ret
	}
	return *o.SecuritySystemManagementNumeric
}

// GetSecuritySystemManagementNumericOk returns a tuple with the SecuritySystemManagementNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetSecuritySystemManagementNumericOk() (*int64, bool) {
	if o == nil || IsNil(o.SecuritySystemManagementNumeric) {
		return nil, false
	}
	return o.SecuritySystemManagementNumeric, true
}

// HasSecuritySystemManagementNumeric returns a boolean if a field has been set.
func (o *SystemResourceInner) HasSecuritySystemManagementNumeric() bool {
	if o != nil && !IsNil(o.SecuritySystemManagementNumeric) {
		return true
	}

	return false
}

// SetSecuritySystemManagementNumeric gets a reference to the given int64 and assigns it to the SecuritySystemManagementNumeric field.
func (o *SystemResourceInner) SetSecuritySystemManagementNumeric(v int64) {
	o.SecuritySystemManagementNumeric = &v
}

// GetSupportedLocales returns the SupportedLocales field value if set, zero value otherwise.
func (o *SystemResourceInner) GetSupportedLocales() string {
	if o == nil || IsNil(o.SupportedLocales) {
		var ret string
		return ret
	}
	return *o.SupportedLocales
}

// GetSupportedLocalesOk returns a tuple with the SupportedLocales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetSupportedLocalesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedLocales) {
		return nil, false
	}
	return o.SupportedLocales, true
}

// HasSupportedLocales returns a boolean if a field has been set.
func (o *SystemResourceInner) HasSupportedLocales() bool {
	if o != nil && !IsNil(o.SupportedLocales) {
		return true
	}

	return false
}

// SetSupportedLocales gets a reference to the given string and assigns it to the SupportedLocales field.
func (o *SystemResourceInner) SetSupportedLocales(v string) {
	o.SupportedLocales = &v
}

// GetSystemContact returns the SystemContact field value if set, zero value otherwise.
func (o *SystemResourceInner) GetSystemContact() string {
	if o == nil || IsNil(o.SystemContact) {
		var ret string
		return ret
	}
	return *o.SystemContact
}

// GetSystemContactOk returns a tuple with the SystemContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetSystemContactOk() (*string, bool) {
	if o == nil || IsNil(o.SystemContact) {
		return nil, false
	}
	return o.SystemContact, true
}

// HasSystemContact returns a boolean if a field has been set.
func (o *SystemResourceInner) HasSystemContact() bool {
	if o != nil && !IsNil(o.SystemContact) {
		return true
	}

	return false
}

// SetSystemContact gets a reference to the given string and assigns it to the SystemContact field.
func (o *SystemResourceInner) SetSystemContact(v string) {
	o.SystemContact = &v
}

// GetSystemInformation returns the SystemInformation field value if set, zero value otherwise.
func (o *SystemResourceInner) GetSystemInformation() string {
	if o == nil || IsNil(o.SystemInformation) {
		var ret string
		return ret
	}
	return *o.SystemInformation
}

// GetSystemInformationOk returns a tuple with the SystemInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetSystemInformationOk() (*string, bool) {
	if o == nil || IsNil(o.SystemInformation) {
		return nil, false
	}
	return o.SystemInformation, true
}

// HasSystemInformation returns a boolean if a field has been set.
func (o *SystemResourceInner) HasSystemInformation() bool {
	if o != nil && !IsNil(o.SystemInformation) {
		return true
	}

	return false
}

// SetSystemInformation gets a reference to the given string and assigns it to the SystemInformation field.
func (o *SystemResourceInner) SetSystemInformation(v string) {
	o.SystemInformation = &v
}

// GetSystemLocation returns the SystemLocation field value if set, zero value otherwise.
func (o *SystemResourceInner) GetSystemLocation() string {
	if o == nil || IsNil(o.SystemLocation) {
		var ret string
		return ret
	}
	return *o.SystemLocation
}

// GetSystemLocationOk returns a tuple with the SystemLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetSystemLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SystemLocation) {
		return nil, false
	}
	return o.SystemLocation, true
}

// HasSystemLocation returns a boolean if a field has been set.
func (o *SystemResourceInner) HasSystemLocation() bool {
	if o != nil && !IsNil(o.SystemLocation) {
		return true
	}

	return false
}

// SetSystemLocation gets a reference to the given string and assigns it to the SystemLocation field.
func (o *SystemResourceInner) SetSystemLocation(v string) {
	o.SystemLocation = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *SystemResourceInner) GetSystemName() string {
	if o == nil || IsNil(o.SystemName) {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetSystemNameOk() (*string, bool) {
	if o == nil || IsNil(o.SystemName) {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *SystemResourceInner) HasSystemName() bool {
	if o != nil && !IsNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *SystemResourceInner) SetSystemName(v string) {
	o.SystemName = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SystemResourceInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SystemResourceInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SystemResourceInner) SetUrl(v string) {
	o.Url = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *SystemResourceInner) GetVendorName() string {
	if o == nil || IsNil(o.VendorName) {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetVendorNameOk() (*string, bool) {
	if o == nil || IsNil(o.VendorName) {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *SystemResourceInner) HasVendorName() bool {
	if o != nil && !IsNil(o.VendorName) {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *SystemResourceInner) SetVendorName(v string) {
	o.VendorName = &v
}

// GetRedundancy returns the Redundancy field value if set, zero value otherwise.
func (o *SystemResourceInner) GetRedundancy() []RedundancyResourceInner {
	if o == nil || IsNil(o.Redundancy) {
		var ret []RedundancyResourceInner
		return ret
	}
	return o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceInner) GetRedundancyOk() ([]RedundancyResourceInner, bool) {
	if o == nil || IsNil(o.Redundancy) {
		return nil, false
	}
	return o.Redundancy, true
}

// HasRedundancy returns a boolean if a field has been set.
func (o *SystemResourceInner) HasRedundancy() bool {
	if o != nil && !IsNil(o.Redundancy) {
		return true
	}

	return false
}

// SetRedundancy gets a reference to the given []RedundancyResourceInner and assigns it to the Redundancy field.
func (o *SystemResourceInner) SetRedundancy(v []RedundancyResourceInner) {
	o.Redundancy = v
}

func (o SystemResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemResourceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectName) {
		toSerialize["object-name"] = o.ObjectName
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.ConfigurationSelector) {
		toSerialize["configuration-selector"] = o.ConfigurationSelector
	}
	if !IsNil(o.CurrentNodeWwn) {
		toSerialize["current-node-wwn"] = o.CurrentNodeWwn
	}
	if !IsNil(o.EnclosureCount) {
		toSerialize["enclosure-count"] = o.EnclosureCount
	}
	if !IsNil(o.FdeSecurityStatus) {
		toSerialize["fde-security-status"] = o.FdeSecurityStatus
	}
	if !IsNil(o.FdeSecurityStatusNumeric) {
		toSerialize["fde-security-status-numeric"] = o.FdeSecurityStatusNumeric
	}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !IsNil(o.HealthNumeric) {
		toSerialize["health-numeric"] = o.HealthNumeric
	}
	if !IsNil(o.HealthReason) {
		toSerialize["health-reason"] = o.HealthReason
	}
	if !IsNil(o.MidplaneSerialNumber) {
		toSerialize["midplane-serial-number"] = o.MidplaneSerialNumber
	}
	if !IsNil(o.OtherMCStatus) {
		toSerialize["other-MC-status"] = o.OtherMCStatus
	}
	if !IsNil(o.OtherMCStatusNumeric) {
		toSerialize["other-MC-status-numeric"] = o.OtherMCStatusNumeric
	}
	if !IsNil(o.PfuStatus) {
		toSerialize["pfuStatus"] = o.PfuStatus
	}
	if !IsNil(o.PfuStatusNumeric) {
		toSerialize["pfuStatus-numeric"] = o.PfuStatusNumeric
	}
	if !IsNil(o.PlatformBrand) {
		toSerialize["platform-brand"] = o.PlatformBrand
	}
	if !IsNil(o.PlatformBrandNumeric) {
		toSerialize["platform-brand-numeric"] = o.PlatformBrandNumeric
	}
	if !IsNil(o.PlatformType) {
		toSerialize["platform-type"] = o.PlatformType
	}
	if !IsNil(o.PlatformTypeNumeric) {
		toSerialize["platform-type-numeric"] = o.PlatformTypeNumeric
	}
	if !IsNil(o.ProductBrand) {
		toSerialize["product-brand"] = o.ProductBrand
	}
	if !IsNil(o.ProductId) {
		toSerialize["product-id"] = o.ProductId
	}
	if !IsNil(o.ScsiProductId) {
		toSerialize["scsi-product-id"] = o.ScsiProductId
	}
	if !IsNil(o.ScsiVendorId) {
		toSerialize["scsi-vendor-id"] = o.ScsiVendorId
	}
	if !IsNil(o.SecuritySystemManagement) {
		toSerialize["security-system-management"] = o.SecuritySystemManagement
	}
	if !IsNil(o.SecuritySystemManagementNumeric) {
		toSerialize["security-system-management-numeric"] = o.SecuritySystemManagementNumeric
	}
	if !IsNil(o.SupportedLocales) {
		toSerialize["supported-locales"] = o.SupportedLocales
	}
	if !IsNil(o.SystemContact) {
		toSerialize["system-contact"] = o.SystemContact
	}
	if !IsNil(o.SystemInformation) {
		toSerialize["system-information"] = o.SystemInformation
	}
	if !IsNil(o.SystemLocation) {
		toSerialize["system-location"] = o.SystemLocation
	}
	if !IsNil(o.SystemName) {
		toSerialize["system-name"] = o.SystemName
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.VendorName) {
		toSerialize["vendor-name"] = o.VendorName
	}
	if !IsNil(o.Redundancy) {
		toSerialize["redundancy"] = o.Redundancy
	}
	return toSerialize, nil
}

type NullableSystemResourceInner struct {
	value *SystemResourceInner
	isSet bool
}

func (v NullableSystemResourceInner) Get() *SystemResourceInner {
	return v.value
}

func (v *NullableSystemResourceInner) Set(val *SystemResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemResourceInner(val *SystemResourceInner) *NullableSystemResourceInner {
	return &NullableSystemResourceInner{value: val, isSet: true}
}

func (v NullableSystemResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
