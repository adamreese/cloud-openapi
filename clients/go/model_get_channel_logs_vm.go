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

// checks if the GetChannelLogsVm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetChannelLogsVm{}

// GetChannelLogsVm struct for GetChannelLogsVm
type GetChannelLogsVm struct {
	Logs []string `json:"logs"`
}

// NewGetChannelLogsVm instantiates a new GetChannelLogsVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetChannelLogsVm(logs []string) *GetChannelLogsVm {
	this := GetChannelLogsVm{}
	this.Logs = logs
	return &this
}

// NewGetChannelLogsVmWithDefaults instantiates a new GetChannelLogsVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetChannelLogsVmWithDefaults() *GetChannelLogsVm {
	this := GetChannelLogsVm{}
	return &this
}

// GetLogs returns the Logs field value
func (o *GetChannelLogsVm) GetLogs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *GetChannelLogsVm) GetLogsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logs, true
}

// SetLogs sets field value
func (o *GetChannelLogsVm) SetLogs(v []string) {
	o.Logs = v
}

func (o GetChannelLogsVm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetChannelLogsVm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logs"] = o.Logs
	return toSerialize, nil
}

type NullableGetChannelLogsVm struct {
	value *GetChannelLogsVm
	isSet bool
}

func (v NullableGetChannelLogsVm) Get() *GetChannelLogsVm {
	return v.value
}

func (v *NullableGetChannelLogsVm) Set(val *GetChannelLogsVm) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChannelLogsVm) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChannelLogsVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChannelLogsVm(val *GetChannelLogsVm) *NullableGetChannelLogsVm {
	return &NullableGetChannelLogsVm{value: val, isSet: true}
}

func (v NullableGetChannelLogsVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChannelLogsVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

