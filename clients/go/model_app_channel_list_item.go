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

// checks if the AppChannelListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppChannelListItem{}

// AppChannelListItem struct for AppChannelListItem
type AppChannelListItem struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ActiveRevisionNumber *string `json:"activeRevisionNumber,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Created time.Time `json:"created"`
	HealthStatus *ApiHealthStatus `json:"healthStatus,omitempty"`
}

// NewAppChannelListItem instantiates a new AppChannelListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppChannelListItem(id string, name string, created time.Time) *AppChannelListItem {
	this := AppChannelListItem{}
	this.Id = id
	this.Name = name
	this.Created = created
	return &this
}

// NewAppChannelListItemWithDefaults instantiates a new AppChannelListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppChannelListItemWithDefaults() *AppChannelListItem {
	this := AppChannelListItem{}
	return &this
}

// GetId returns the Id field value
func (o *AppChannelListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppChannelListItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppChannelListItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AppChannelListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppChannelListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppChannelListItem) SetName(v string) {
	o.Name = v
}

// GetActiveRevisionNumber returns the ActiveRevisionNumber field value if set, zero value otherwise.
func (o *AppChannelListItem) GetActiveRevisionNumber() string {
	if o == nil || IsNil(o.ActiveRevisionNumber) {
		var ret string
		return ret
	}
	return *o.ActiveRevisionNumber
}

// GetActiveRevisionNumberOk returns a tuple with the ActiveRevisionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppChannelListItem) GetActiveRevisionNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ActiveRevisionNumber) {
		return nil, false
	}
	return o.ActiveRevisionNumber, true
}

// HasActiveRevisionNumber returns a boolean if a field has been set.
func (o *AppChannelListItem) HasActiveRevisionNumber() bool {
	if o != nil && !IsNil(o.ActiveRevisionNumber) {
		return true
	}

	return false
}

// SetActiveRevisionNumber gets a reference to the given string and assigns it to the ActiveRevisionNumber field.
func (o *AppChannelListItem) SetActiveRevisionNumber(v string) {
	o.ActiveRevisionNumber = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *AppChannelListItem) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppChannelListItem) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *AppChannelListItem) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *AppChannelListItem) SetDomain(v string) {
	o.Domain = &v
}

// GetCreated returns the Created field value
func (o *AppChannelListItem) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AppChannelListItem) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AppChannelListItem) SetCreated(v time.Time) {
	o.Created = v
}

// GetHealthStatus returns the HealthStatus field value if set, zero value otherwise.
func (o *AppChannelListItem) GetHealthStatus() ApiHealthStatus {
	if o == nil || IsNil(o.HealthStatus) {
		var ret ApiHealthStatus
		return ret
	}
	return *o.HealthStatus
}

// GetHealthStatusOk returns a tuple with the HealthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppChannelListItem) GetHealthStatusOk() (*ApiHealthStatus, bool) {
	if o == nil || IsNil(o.HealthStatus) {
		return nil, false
	}
	return o.HealthStatus, true
}

// HasHealthStatus returns a boolean if a field has been set.
func (o *AppChannelListItem) HasHealthStatus() bool {
	if o != nil && !IsNil(o.HealthStatus) {
		return true
	}

	return false
}

// SetHealthStatus gets a reference to the given ApiHealthStatus and assigns it to the HealthStatus field.
func (o *AppChannelListItem) SetHealthStatus(v ApiHealthStatus) {
	o.HealthStatus = &v
}

func (o AppChannelListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppChannelListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.ActiveRevisionNumber) {
		toSerialize["activeRevisionNumber"] = o.ActiveRevisionNumber
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	toSerialize["created"] = o.Created
	if !IsNil(o.HealthStatus) {
		toSerialize["healthStatus"] = o.HealthStatus
	}
	return toSerialize, nil
}

type NullableAppChannelListItem struct {
	value *AppChannelListItem
	isSet bool
}

func (v NullableAppChannelListItem) Get() *AppChannelListItem {
	return v.value
}

func (v *NullableAppChannelListItem) Set(val *AppChannelListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAppChannelListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAppChannelListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppChannelListItem(val *AppChannelListItem) *NullableAppChannelListItem {
	return &NullableAppChannelListItem{value: val, isSet: true}
}

func (v NullableAppChannelListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppChannelListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

