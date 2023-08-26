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

// checks if the CreateChannelCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChannelCommand{}

// CreateChannelCommand struct for CreateChannelCommand
type CreateChannelCommand struct {
	AppId string `json:"appId"`
	Name string `json:"name"`
	RevisionSelectionStrategy ChannelRevisionSelectionStrategy `json:"revisionSelectionStrategy"`
	RangeRule NullableString `json:"rangeRule,omitempty"`
	ActiveRevisionId NullableString `json:"activeRevisionId,omitempty"`
}

// NewCreateChannelCommand instantiates a new CreateChannelCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChannelCommand(appId string, name string, revisionSelectionStrategy ChannelRevisionSelectionStrategy) *CreateChannelCommand {
	this := CreateChannelCommand{}
	this.AppId = appId
	this.Name = name
	this.RevisionSelectionStrategy = revisionSelectionStrategy
	return &this
}

// NewCreateChannelCommandWithDefaults instantiates a new CreateChannelCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChannelCommandWithDefaults() *CreateChannelCommand {
	this := CreateChannelCommand{}
	return &this
}

// GetAppId returns the AppId field value
func (o *CreateChannelCommand) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *CreateChannelCommand) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *CreateChannelCommand) SetAppId(v string) {
	o.AppId = v
}

// GetName returns the Name field value
func (o *CreateChannelCommand) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateChannelCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateChannelCommand) SetName(v string) {
	o.Name = v
}

// GetRevisionSelectionStrategy returns the RevisionSelectionStrategy field value
func (o *CreateChannelCommand) GetRevisionSelectionStrategy() ChannelRevisionSelectionStrategy {
	if o == nil {
		var ret ChannelRevisionSelectionStrategy
		return ret
	}

	return o.RevisionSelectionStrategy
}

// GetRevisionSelectionStrategyOk returns a tuple with the RevisionSelectionStrategy field value
// and a boolean to check if the value has been set.
func (o *CreateChannelCommand) GetRevisionSelectionStrategyOk() (*ChannelRevisionSelectionStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionSelectionStrategy, true
}

// SetRevisionSelectionStrategy sets field value
func (o *CreateChannelCommand) SetRevisionSelectionStrategy(v ChannelRevisionSelectionStrategy) {
	o.RevisionSelectionStrategy = v
}

// GetRangeRule returns the RangeRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChannelCommand) GetRangeRule() string {
	if o == nil || IsNil(o.RangeRule.Get()) {
		var ret string
		return ret
	}
	return *o.RangeRule.Get()
}

// GetRangeRuleOk returns a tuple with the RangeRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChannelCommand) GetRangeRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RangeRule.Get(), o.RangeRule.IsSet()
}

// HasRangeRule returns a boolean if a field has been set.
func (o *CreateChannelCommand) HasRangeRule() bool {
	if o != nil && o.RangeRule.IsSet() {
		return true
	}

	return false
}

// SetRangeRule gets a reference to the given NullableString and assigns it to the RangeRule field.
func (o *CreateChannelCommand) SetRangeRule(v string) {
	o.RangeRule.Set(&v)
}
// SetRangeRuleNil sets the value for RangeRule to be an explicit nil
func (o *CreateChannelCommand) SetRangeRuleNil() {
	o.RangeRule.Set(nil)
}

// UnsetRangeRule ensures that no value is present for RangeRule, not even an explicit nil
func (o *CreateChannelCommand) UnsetRangeRule() {
	o.RangeRule.Unset()
}

// GetActiveRevisionId returns the ActiveRevisionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChannelCommand) GetActiveRevisionId() string {
	if o == nil || IsNil(o.ActiveRevisionId.Get()) {
		var ret string
		return ret
	}
	return *o.ActiveRevisionId.Get()
}

// GetActiveRevisionIdOk returns a tuple with the ActiveRevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChannelCommand) GetActiveRevisionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveRevisionId.Get(), o.ActiveRevisionId.IsSet()
}

// HasActiveRevisionId returns a boolean if a field has been set.
func (o *CreateChannelCommand) HasActiveRevisionId() bool {
	if o != nil && o.ActiveRevisionId.IsSet() {
		return true
	}

	return false
}

// SetActiveRevisionId gets a reference to the given NullableString and assigns it to the ActiveRevisionId field.
func (o *CreateChannelCommand) SetActiveRevisionId(v string) {
	o.ActiveRevisionId.Set(&v)
}
// SetActiveRevisionIdNil sets the value for ActiveRevisionId to be an explicit nil
func (o *CreateChannelCommand) SetActiveRevisionIdNil() {
	o.ActiveRevisionId.Set(nil)
}

// UnsetActiveRevisionId ensures that no value is present for ActiveRevisionId, not even an explicit nil
func (o *CreateChannelCommand) UnsetActiveRevisionId() {
	o.ActiveRevisionId.Unset()
}

func (o CreateChannelCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChannelCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	toSerialize["name"] = o.Name
	toSerialize["revisionSelectionStrategy"] = o.RevisionSelectionStrategy
	if o.RangeRule.IsSet() {
		toSerialize["rangeRule"] = o.RangeRule.Get()
	}
	if o.ActiveRevisionId.IsSet() {
		toSerialize["activeRevisionId"] = o.ActiveRevisionId.Get()
	}
	return toSerialize, nil
}

type NullableCreateChannelCommand struct {
	value *CreateChannelCommand
	isSet bool
}

func (v NullableCreateChannelCommand) Get() *CreateChannelCommand {
	return v.value
}

func (v *NullableCreateChannelCommand) Set(val *CreateChannelCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChannelCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChannelCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChannelCommand(val *CreateChannelCommand) *NullableCreateChannelCommand {
	return &NullableCreateChannelCommand{value: val, isSet: true}
}

func (v NullableCreateChannelCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChannelCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


