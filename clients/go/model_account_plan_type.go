/*
Fermyon Cloud API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud_openapi

import (
	"encoding/json"
	"fmt"
)

// AccountPlanType the model 'AccountPlanType'
type AccountPlanType string

// List of AccountPlanType
const (
	FREE AccountPlanType = "Free"
	GROWTH AccountPlanType = "Growth"
	ENTERPRISE AccountPlanType = "Enterprise"
)

// All allowed values of AccountPlanType enum
var AllowedAccountPlanTypeEnumValues = []AccountPlanType{
	"Free",
	"Growth",
	"Enterprise",
}

func (v *AccountPlanType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountPlanType(value)
	for _, existing := range AllowedAccountPlanTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountPlanType", value)
}

// NewAccountPlanTypeFromValue returns a pointer to a valid AccountPlanType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountPlanTypeFromValue(v string) (*AccountPlanType, error) {
	ev := AccountPlanType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountPlanType: valid values are %v", v, AllowedAccountPlanTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountPlanType) IsValid() bool {
	for _, existing := range AllowedAccountPlanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountPlanType value
func (v AccountPlanType) Ptr() *AccountPlanType {
	return &v
}

type NullableAccountPlanType struct {
	value *AccountPlanType
	isSet bool
}

func (v NullableAccountPlanType) Get() *AccountPlanType {
	return v.value
}

func (v *NullableAccountPlanType) Set(val *AccountPlanType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPlanType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPlanType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPlanType(val *AccountPlanType) *NullableAccountPlanType {
	return &NullableAccountPlanType{value: val, isSet: true}
}

func (v NullableAccountPlanType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPlanType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

