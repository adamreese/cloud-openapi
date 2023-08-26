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

// DesiredStatus the model 'DesiredStatus'
type DesiredStatus string

// List of DesiredStatus
const (
	RUNNING DesiredStatus = "Running"
	DEAD DesiredStatus = "Dead"
)

// All allowed values of DesiredStatus enum
var AllowedDesiredStatusEnumValues = []DesiredStatus{
	"Running",
	"Dead",
}

func (v *DesiredStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DesiredStatus(value)
	for _, existing := range AllowedDesiredStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DesiredStatus", value)
}

// NewDesiredStatusFromValue returns a pointer to a valid DesiredStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDesiredStatusFromValue(v string) (*DesiredStatus, error) {
	ev := DesiredStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DesiredStatus: valid values are %v", v, AllowedDesiredStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DesiredStatus) IsValid() bool {
	for _, existing := range AllowedDesiredStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DesiredStatus value
func (v DesiredStatus) Ptr() *DesiredStatus {
	return &v
}

type NullableDesiredStatus struct {
	value *DesiredStatus
	isSet bool
}

func (v NullableDesiredStatus) Get() *DesiredStatus {
	return v.value
}

func (v *NullableDesiredStatus) Set(val *DesiredStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDesiredStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDesiredStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesiredStatus(val *DesiredStatus) *NullableDesiredStatus {
	return &NullableDesiredStatus{value: val, isSet: true}
}

func (v NullableDesiredStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesiredStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

