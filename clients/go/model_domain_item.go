/*
Fermyon Cloud API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud_openapi

import (
	"encoding/json"
	"time"
)

// checks if the DomainItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainItem{}

// DomainItem struct for DomainItem
type DomainItem struct {
	Name string `json:"name"`
	ValidationStatus string `json:"validationStatus"`
	ValidatedAt NullableTime `json:"validatedAt,omitempty"`
	DnsRecords []DnsRecord `json:"dnsRecords"`
}

// NewDomainItem instantiates a new DomainItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainItem(name string, validationStatus string, dnsRecords []DnsRecord) *DomainItem {
	this := DomainItem{}
	this.Name = name
	this.ValidationStatus = validationStatus
	this.DnsRecords = dnsRecords
	return &this
}

// NewDomainItemWithDefaults instantiates a new DomainItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainItemWithDefaults() *DomainItem {
	this := DomainItem{}
	return &this
}

// GetName returns the Name field value
func (o *DomainItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DomainItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DomainItem) SetName(v string) {
	o.Name = v
}

// GetValidationStatus returns the ValidationStatus field value
func (o *DomainItem) GetValidationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value
// and a boolean to check if the value has been set.
func (o *DomainItem) GetValidationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidationStatus, true
}

// SetValidationStatus sets field value
func (o *DomainItem) SetValidationStatus(v string) {
	o.ValidationStatus = v
}

// GetValidatedAt returns the ValidatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainItem) GetValidatedAt() time.Time {
	if o == nil || IsNil(o.ValidatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ValidatedAt.Get()
}

// GetValidatedAtOk returns a tuple with the ValidatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainItem) GetValidatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidatedAt.Get(), o.ValidatedAt.IsSet()
}

// HasValidatedAt returns a boolean if a field has been set.
func (o *DomainItem) HasValidatedAt() bool {
	if o != nil && o.ValidatedAt.IsSet() {
		return true
	}

	return false
}

// SetValidatedAt gets a reference to the given NullableTime and assigns it to the ValidatedAt field.
func (o *DomainItem) SetValidatedAt(v time.Time) {
	o.ValidatedAt.Set(&v)
}
// SetValidatedAtNil sets the value for ValidatedAt to be an explicit nil
func (o *DomainItem) SetValidatedAtNil() {
	o.ValidatedAt.Set(nil)
}

// UnsetValidatedAt ensures that no value is present for ValidatedAt, not even an explicit nil
func (o *DomainItem) UnsetValidatedAt() {
	o.ValidatedAt.Unset()
}

// GetDnsRecords returns the DnsRecords field value
func (o *DomainItem) GetDnsRecords() []DnsRecord {
	if o == nil {
		var ret []DnsRecord
		return ret
	}

	return o.DnsRecords
}

// GetDnsRecordsOk returns a tuple with the DnsRecords field value
// and a boolean to check if the value has been set.
func (o *DomainItem) GetDnsRecordsOk() ([]DnsRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsRecords, true
}

// SetDnsRecords sets field value
func (o *DomainItem) SetDnsRecords(v []DnsRecord) {
	o.DnsRecords = v
}

func (o DomainItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["validationStatus"] = o.ValidationStatus
	if o.ValidatedAt.IsSet() {
		toSerialize["validatedAt"] = o.ValidatedAt.Get()
	}
	toSerialize["dnsRecords"] = o.DnsRecords
	return toSerialize, nil
}

type NullableDomainItem struct {
	value *DomainItem
	isSet bool
}

func (v NullableDomainItem) Get() *DomainItem {
	return v.value
}

func (v *NullableDomainItem) Set(val *DomainItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainItem(val *DomainItem) *NullableDomainItem {
	return &NullableDomainItem{value: val, isSet: true}
}

func (v NullableDomainItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

