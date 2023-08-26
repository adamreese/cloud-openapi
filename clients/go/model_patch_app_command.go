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

// checks if the PatchAppCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchAppCommand{}

// PatchAppCommand struct for PatchAppCommand
type PatchAppCommand struct {
	Id string `json:"id"`
	Domain *StringField `json:"domain,omitempty"`
}

// NewPatchAppCommand instantiates a new PatchAppCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchAppCommand(id string) *PatchAppCommand {
	this := PatchAppCommand{}
	this.Id = id
	return &this
}

// NewPatchAppCommandWithDefaults instantiates a new PatchAppCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchAppCommandWithDefaults() *PatchAppCommand {
	this := PatchAppCommand{}
	return &this
}

// GetId returns the Id field value
func (o *PatchAppCommand) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchAppCommand) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchAppCommand) SetId(v string) {
	o.Id = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *PatchAppCommand) GetDomain() StringField {
	if o == nil || IsNil(o.Domain) {
		var ret StringField
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAppCommand) GetDomainOk() (*StringField, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *PatchAppCommand) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given StringField and assigns it to the Domain field.
func (o *PatchAppCommand) SetDomain(v StringField) {
	o.Domain = &v
}

func (o PatchAppCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchAppCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	return toSerialize, nil
}

type NullablePatchAppCommand struct {
	value *PatchAppCommand
	isSet bool
}

func (v NullablePatchAppCommand) Get() *PatchAppCommand {
	return v.value
}

func (v *NullablePatchAppCommand) Set(val *PatchAppCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchAppCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchAppCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchAppCommand(val *PatchAppCommand) *NullablePatchAppCommand {
	return &NullablePatchAppCommand{value: val, isSet: true}
}

func (v NullablePatchAppCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchAppCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


