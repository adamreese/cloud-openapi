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

// checks if the StringPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringPage{}

// StringPage struct for StringPage
type StringPage struct {
	Items []string `json:"items"`
	TotalItems int32 `json:"totalItems"`
	PageIndex int32 `json:"pageIndex"`
	PageSize int32 `json:"pageSize"`
	IsLastPage bool `json:"isLastPage"`
}

// NewStringPage instantiates a new StringPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringPage(items []string, totalItems int32, pageIndex int32, pageSize int32, isLastPage bool) *StringPage {
	this := StringPage{}
	this.Items = items
	this.TotalItems = totalItems
	this.PageIndex = pageIndex
	this.PageSize = pageSize
	this.IsLastPage = isLastPage
	return &this
}

// NewStringPageWithDefaults instantiates a new StringPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringPageWithDefaults() *StringPage {
	this := StringPage{}
	return &this
}

// GetItems returns the Items field value
func (o *StringPage) GetItems() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *StringPage) GetItemsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *StringPage) SetItems(v []string) {
	o.Items = v
}

// GetTotalItems returns the TotalItems field value
func (o *StringPage) GetTotalItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value
// and a boolean to check if the value has been set.
func (o *StringPage) GetTotalItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalItems, true
}

// SetTotalItems sets field value
func (o *StringPage) SetTotalItems(v int32) {
	o.TotalItems = v
}

// GetPageIndex returns the PageIndex field value
func (o *StringPage) GetPageIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value
// and a boolean to check if the value has been set.
func (o *StringPage) GetPageIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageIndex, true
}

// SetPageIndex sets field value
func (o *StringPage) SetPageIndex(v int32) {
	o.PageIndex = v
}

// GetPageSize returns the PageSize field value
func (o *StringPage) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *StringPage) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *StringPage) SetPageSize(v int32) {
	o.PageSize = v
}

// GetIsLastPage returns the IsLastPage field value
func (o *StringPage) GetIsLastPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLastPage
}

// GetIsLastPageOk returns a tuple with the IsLastPage field value
// and a boolean to check if the value has been set.
func (o *StringPage) GetIsLastPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLastPage, true
}

// SetIsLastPage sets field value
func (o *StringPage) SetIsLastPage(v bool) {
	o.IsLastPage = v
}

func (o StringPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["totalItems"] = o.TotalItems
	toSerialize["pageIndex"] = o.PageIndex
	toSerialize["pageSize"] = o.PageSize
	// skip: isLastPage is readOnly
	return toSerialize, nil
}

type NullableStringPage struct {
	value *StringPage
	isSet bool
}

func (v NullableStringPage) Get() *StringPage {
	return v.value
}

func (v *NullableStringPage) Set(val *StringPage) {
	v.value = val
	v.isSet = true
}

func (v NullableStringPage) IsSet() bool {
	return v.isSet
}

func (v *NullableStringPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringPage(val *StringPage) *NullableStringPage {
	return &NullableStringPage{value: val, isSet: true}
}

func (v NullableStringPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


