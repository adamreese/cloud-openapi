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

// checks if the CreateVariablePairCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVariablePairCommand{}

// CreateVariablePairCommand struct for CreateVariablePairCommand
type CreateVariablePairCommand struct {
	AppId string `json:"appId"`
	Variable string `json:"variable"`
	Value string `json:"value"`
}

// NewCreateVariablePairCommand instantiates a new CreateVariablePairCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVariablePairCommand(appId string, variable string, value string) *CreateVariablePairCommand {
	this := CreateVariablePairCommand{}
	this.AppId = appId
	this.Variable = variable
	this.Value = value
	return &this
}

// NewCreateVariablePairCommandWithDefaults instantiates a new CreateVariablePairCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVariablePairCommandWithDefaults() *CreateVariablePairCommand {
	this := CreateVariablePairCommand{}
	return &this
}

// GetAppId returns the AppId field value
func (o *CreateVariablePairCommand) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *CreateVariablePairCommand) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *CreateVariablePairCommand) SetAppId(v string) {
	o.AppId = v
}

// GetVariable returns the Variable field value
func (o *CreateVariablePairCommand) GetVariable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Variable
}

// GetVariableOk returns a tuple with the Variable field value
// and a boolean to check if the value has been set.
func (o *CreateVariablePairCommand) GetVariableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variable, true
}

// SetVariable sets field value
func (o *CreateVariablePairCommand) SetVariable(v string) {
	o.Variable = v
}

// GetValue returns the Value field value
func (o *CreateVariablePairCommand) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreateVariablePairCommand) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreateVariablePairCommand) SetValue(v string) {
	o.Value = v
}

func (o CreateVariablePairCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVariablePairCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	toSerialize["variable"] = o.Variable
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableCreateVariablePairCommand struct {
	value *CreateVariablePairCommand
	isSet bool
}

func (v NullableCreateVariablePairCommand) Get() *CreateVariablePairCommand {
	return v.value
}

func (v *NullableCreateVariablePairCommand) Set(val *CreateVariablePairCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVariablePairCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVariablePairCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVariablePairCommand(val *CreateVariablePairCommand) *NullableCreateVariablePairCommand {
	return &NullableCreateVariablePairCommand{value: val, isSet: true}
}

func (v NullableCreateVariablePairCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVariablePairCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

