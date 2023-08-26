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

// checks if the CreateSqlDatabaseCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSqlDatabaseCommand{}

// CreateSqlDatabaseCommand struct for CreateSqlDatabaseCommand
type CreateSqlDatabaseCommand struct {
	Name string `json:"name"`
	AppId NullableString `json:"appId,omitempty"`
}

// NewCreateSqlDatabaseCommand instantiates a new CreateSqlDatabaseCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSqlDatabaseCommand(name string) *CreateSqlDatabaseCommand {
	this := CreateSqlDatabaseCommand{}
	this.Name = name
	return &this
}

// NewCreateSqlDatabaseCommandWithDefaults instantiates a new CreateSqlDatabaseCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSqlDatabaseCommandWithDefaults() *CreateSqlDatabaseCommand {
	this := CreateSqlDatabaseCommand{}
	return &this
}

// GetName returns the Name field value
func (o *CreateSqlDatabaseCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSqlDatabaseCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSqlDatabaseCommand) SetName(v string) {
	o.Name = v
}

// GetAppId returns the AppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSqlDatabaseCommand) GetAppId() string {
	if o == nil || IsNil(o.AppId.Get()) {
		var ret string
		return ret
	}
	return *o.AppId.Get()
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSqlDatabaseCommand) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppId.Get(), o.AppId.IsSet()
}

// HasAppId returns a boolean if a field has been set.
func (o *CreateSqlDatabaseCommand) HasAppId() bool {
	if o != nil && o.AppId.IsSet() {
		return true
	}

	return false
}

// SetAppId gets a reference to the given NullableString and assigns it to the AppId field.
func (o *CreateSqlDatabaseCommand) SetAppId(v string) {
	o.AppId.Set(&v)
}
// SetAppIdNil sets the value for AppId to be an explicit nil
func (o *CreateSqlDatabaseCommand) SetAppIdNil() {
	o.AppId.Set(nil)
}

// UnsetAppId ensures that no value is present for AppId, not even an explicit nil
func (o *CreateSqlDatabaseCommand) UnsetAppId() {
	o.AppId.Unset()
}

func (o CreateSqlDatabaseCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSqlDatabaseCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.AppId.IsSet() {
		toSerialize["appId"] = o.AppId.Get()
	}
	return toSerialize, nil
}

type NullableCreateSqlDatabaseCommand struct {
	value *CreateSqlDatabaseCommand
	isSet bool
}

func (v NullableCreateSqlDatabaseCommand) Get() *CreateSqlDatabaseCommand {
	return v.value
}

func (v *NullableCreateSqlDatabaseCommand) Set(val *CreateSqlDatabaseCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSqlDatabaseCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSqlDatabaseCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSqlDatabaseCommand(val *CreateSqlDatabaseCommand) *NullableCreateSqlDatabaseCommand {
	return &NullableCreateSqlDatabaseCommand{value: val, isSet: true}
}

func (v NullableCreateSqlDatabaseCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSqlDatabaseCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

