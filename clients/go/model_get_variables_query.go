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

// checks if the GetVariablesQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVariablesQuery{}

// GetVariablesQuery struct for GetVariablesQuery
type GetVariablesQuery struct {
	AppId string `json:"appId"`
}

// NewGetVariablesQuery instantiates a new GetVariablesQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVariablesQuery(appId string) *GetVariablesQuery {
	this := GetVariablesQuery{}
	this.AppId = appId
	return &this
}

// NewGetVariablesQueryWithDefaults instantiates a new GetVariablesQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVariablesQueryWithDefaults() *GetVariablesQuery {
	this := GetVariablesQuery{}
	return &this
}

// GetAppId returns the AppId field value
func (o *GetVariablesQuery) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *GetVariablesQuery) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *GetVariablesQuery) SetAppId(v string) {
	o.AppId = v
}

func (o GetVariablesQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVariablesQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	return toSerialize, nil
}

type NullableGetVariablesQuery struct {
	value *GetVariablesQuery
	isSet bool
}

func (v NullableGetVariablesQuery) Get() *GetVariablesQuery {
	return v.value
}

func (v *NullableGetVariablesQuery) Set(val *GetVariablesQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVariablesQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVariablesQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVariablesQuery(val *GetVariablesQuery) *NullableGetVariablesQuery {
	return &NullableGetVariablesQuery{value: val, isSet: true}
}

func (v NullableGetVariablesQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVariablesQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

