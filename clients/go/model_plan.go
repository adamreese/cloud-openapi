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

// checks if the Plan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Plan{}

// Plan struct for Plan
type Plan struct {
	AccountPlanType *AccountPlanType `json:"accountPlanType,omitempty"`
	Price NullableFloat64 `json:"price,omitempty"`
}

// NewPlan instantiates a new Plan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlan() *Plan {
	this := Plan{}
	return &this
}

// NewPlanWithDefaults instantiates a new Plan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanWithDefaults() *Plan {
	this := Plan{}
	return &this
}

// GetAccountPlanType returns the AccountPlanType field value if set, zero value otherwise.
func (o *Plan) GetAccountPlanType() AccountPlanType {
	if o == nil || IsNil(o.AccountPlanType) {
		var ret AccountPlanType
		return ret
	}
	return *o.AccountPlanType
}

// GetAccountPlanTypeOk returns a tuple with the AccountPlanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetAccountPlanTypeOk() (*AccountPlanType, bool) {
	if o == nil || IsNil(o.AccountPlanType) {
		return nil, false
	}
	return o.AccountPlanType, true
}

// HasAccountPlanType returns a boolean if a field has been set.
func (o *Plan) HasAccountPlanType() bool {
	if o != nil && !IsNil(o.AccountPlanType) {
		return true
	}

	return false
}

// SetAccountPlanType gets a reference to the given AccountPlanType and assigns it to the AccountPlanType field.
func (o *Plan) SetAccountPlanType(v AccountPlanType) {
	o.AccountPlanType = &v
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Plan) GetPrice() float64 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float64
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Plan) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *Plan) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat64 and assigns it to the Price field.
func (o *Plan) SetPrice(v float64) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *Plan) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *Plan) UnsetPrice() {
	o.Price.Unset()
}

func (o Plan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Plan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountPlanType) {
		toSerialize["accountPlanType"] = o.AccountPlanType
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	return toSerialize, nil
}

type NullablePlan struct {
	value *Plan
	isSet bool
}

func (v NullablePlan) Get() *Plan {
	return v.value
}

func (v *NullablePlan) Set(val *Plan) {
	v.value = val
	v.isSet = true
}

func (v NullablePlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlan(val *Plan) *NullablePlan {
	return &NullablePlan{value: val, isSet: true}
}

func (v NullablePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


