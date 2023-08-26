/*
Fermyon Cloud API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud_openapi

import (
	"encoding/json"
)

// checks if the GuidNullableField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuidNullableField{}

// GuidNullableField struct for GuidNullableField
type GuidNullableField struct {
	Value NullableString `json:"value,omitempty"`
}

// NewGuidNullableField instantiates a new GuidNullableField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuidNullableField() *GuidNullableField {
	this := GuidNullableField{}
	return &this
}

// NewGuidNullableFieldWithDefaults instantiates a new GuidNullableField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuidNullableFieldWithDefaults() *GuidNullableField {
	this := GuidNullableField{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidNullableField) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidNullableField) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *GuidNullableField) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *GuidNullableField) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *GuidNullableField) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *GuidNullableField) UnsetValue() {
	o.Value.Unset()
}

func (o GuidNullableField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuidNullableField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableGuidNullableField struct {
	value *GuidNullableField
	isSet bool
}

func (v NullableGuidNullableField) Get() *GuidNullableField {
	return v.value
}

func (v *NullableGuidNullableField) Set(val *GuidNullableField) {
	v.value = val
	v.isSet = true
}

func (v NullableGuidNullableField) IsSet() bool {
	return v.isSet
}

func (v *NullableGuidNullableField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuidNullableField(val *GuidNullableField) *NullableGuidNullableField {
	return &NullableGuidNullableField{value: val, isSet: true}
}

func (v NullableGuidNullableField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuidNullableField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


