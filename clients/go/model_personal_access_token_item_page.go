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

// checks if the PersonalAccessTokenItemPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonalAccessTokenItemPage{}

// PersonalAccessTokenItemPage struct for PersonalAccessTokenItemPage
type PersonalAccessTokenItemPage struct {
	Items []PersonalAccessTokenItem `json:"items"`
	TotalItems int32 `json:"totalItems"`
	PageIndex int32 `json:"pageIndex"`
	PageSize int32 `json:"pageSize"`
	IsLastPage bool `json:"isLastPage"`
}

// NewPersonalAccessTokenItemPage instantiates a new PersonalAccessTokenItemPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonalAccessTokenItemPage(items []PersonalAccessTokenItem, totalItems int32, pageIndex int32, pageSize int32, isLastPage bool) *PersonalAccessTokenItemPage {
	this := PersonalAccessTokenItemPage{}
	this.Items = items
	this.TotalItems = totalItems
	this.PageIndex = pageIndex
	this.PageSize = pageSize
	this.IsLastPage = isLastPage
	return &this
}

// NewPersonalAccessTokenItemPageWithDefaults instantiates a new PersonalAccessTokenItemPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonalAccessTokenItemPageWithDefaults() *PersonalAccessTokenItemPage {
	this := PersonalAccessTokenItemPage{}
	return &this
}

// GetItems returns the Items field value
func (o *PersonalAccessTokenItemPage) GetItems() []PersonalAccessTokenItem {
	if o == nil {
		var ret []PersonalAccessTokenItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenItemPage) GetItemsOk() ([]PersonalAccessTokenItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PersonalAccessTokenItemPage) SetItems(v []PersonalAccessTokenItem) {
	o.Items = v
}

// GetTotalItems returns the TotalItems field value
func (o *PersonalAccessTokenItemPage) GetTotalItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenItemPage) GetTotalItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalItems, true
}

// SetTotalItems sets field value
func (o *PersonalAccessTokenItemPage) SetTotalItems(v int32) {
	o.TotalItems = v
}

// GetPageIndex returns the PageIndex field value
func (o *PersonalAccessTokenItemPage) GetPageIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenItemPage) GetPageIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageIndex, true
}

// SetPageIndex sets field value
func (o *PersonalAccessTokenItemPage) SetPageIndex(v int32) {
	o.PageIndex = v
}

// GetPageSize returns the PageSize field value
func (o *PersonalAccessTokenItemPage) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenItemPage) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *PersonalAccessTokenItemPage) SetPageSize(v int32) {
	o.PageSize = v
}

// GetIsLastPage returns the IsLastPage field value
func (o *PersonalAccessTokenItemPage) GetIsLastPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLastPage
}

// GetIsLastPageOk returns a tuple with the IsLastPage field value
// and a boolean to check if the value has been set.
func (o *PersonalAccessTokenItemPage) GetIsLastPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLastPage, true
}

// SetIsLastPage sets field value
func (o *PersonalAccessTokenItemPage) SetIsLastPage(v bool) {
	o.IsLastPage = v
}

func (o PersonalAccessTokenItemPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonalAccessTokenItemPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["totalItems"] = o.TotalItems
	toSerialize["pageIndex"] = o.PageIndex
	toSerialize["pageSize"] = o.PageSize
	// skip: isLastPage is readOnly
	return toSerialize, nil
}

type NullablePersonalAccessTokenItemPage struct {
	value *PersonalAccessTokenItemPage
	isSet bool
}

func (v NullablePersonalAccessTokenItemPage) Get() *PersonalAccessTokenItemPage {
	return v.value
}

func (v *NullablePersonalAccessTokenItemPage) Set(val *PersonalAccessTokenItemPage) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalAccessTokenItemPage) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalAccessTokenItemPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalAccessTokenItemPage(val *PersonalAccessTokenItemPage) *NullablePersonalAccessTokenItemPage {
	return &NullablePersonalAccessTokenItemPage{value: val, isSet: true}
}

func (v NullablePersonalAccessTokenItemPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalAccessTokenItemPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

