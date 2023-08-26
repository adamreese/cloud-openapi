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

// checks if the PersonalAccessTokenItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonalAccessTokenItem{}

// PersonalAccessTokenItem struct for PersonalAccessTokenItem
type PersonalAccessTokenItem struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewPersonalAccessTokenItem instantiates a new PersonalAccessTokenItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonalAccessTokenItem() *PersonalAccessTokenItem {
	this := PersonalAccessTokenItem{}
	return &this
}

// NewPersonalAccessTokenItemWithDefaults instantiates a new PersonalAccessTokenItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonalAccessTokenItemWithDefaults() *PersonalAccessTokenItem {
	this := PersonalAccessTokenItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PersonalAccessTokenItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PersonalAccessTokenItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PersonalAccessTokenItem) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PersonalAccessTokenItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PersonalAccessTokenItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PersonalAccessTokenItem) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PersonalAccessTokenItem) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PersonalAccessTokenItem) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PersonalAccessTokenItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o PersonalAccessTokenItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonalAccessTokenItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullablePersonalAccessTokenItem struct {
	value *PersonalAccessTokenItem
	isSet bool
}

func (v NullablePersonalAccessTokenItem) Get() *PersonalAccessTokenItem {
	return v.value
}

func (v *NullablePersonalAccessTokenItem) Set(val *PersonalAccessTokenItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalAccessTokenItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalAccessTokenItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalAccessTokenItem(val *PersonalAccessTokenItem) *NullablePersonalAccessTokenItem {
	return &NullablePersonalAccessTokenItem{value: val, isSet: true}
}

func (v NullablePersonalAccessTokenItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalAccessTokenItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


