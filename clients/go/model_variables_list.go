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

// checks if the VariablesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariablesList{}

// VariablesList struct for VariablesList
type VariablesList struct {
	Vars []string `json:"vars"`
}

// NewVariablesList instantiates a new VariablesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariablesList(vars []string) *VariablesList {
	this := VariablesList{}
	this.Vars = vars
	return &this
}

// NewVariablesListWithDefaults instantiates a new VariablesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariablesListWithDefaults() *VariablesList {
	this := VariablesList{}
	return &this
}

// GetVars returns the Vars field value
func (o *VariablesList) GetVars() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Vars
}

// GetVarsOk returns a tuple with the Vars field value
// and a boolean to check if the value has been set.
func (o *VariablesList) GetVarsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vars, true
}

// SetVars sets field value
func (o *VariablesList) SetVars(v []string) {
	o.Vars = v
}

func (o VariablesList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariablesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vars"] = o.Vars
	return toSerialize, nil
}

type NullableVariablesList struct {
	value *VariablesList
	isSet bool
}

func (v NullableVariablesList) Get() *VariablesList {
	return v.value
}

func (v *NullableVariablesList) Set(val *VariablesList) {
	v.value = val
	v.isSet = true
}

func (v NullableVariablesList) IsSet() bool {
	return v.isSet
}

func (v *NullableVariablesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariablesList(val *VariablesList) *NullableVariablesList {
	return &NullableVariablesList{value: val, isSet: true}
}

func (v NullableVariablesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariablesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

