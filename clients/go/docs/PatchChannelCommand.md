# PatchChannelCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to [**StringField**](StringField.md) |  | [optional] 
**Domain** | Pointer to [**StringField**](StringField.md) |  | [optional] 
**RevisionSelectionStrategy** | Pointer to [**ChannelRevisionSelectionStrategyField**](ChannelRevisionSelectionStrategyField.md) |  | [optional] 
**RangeRule** | Pointer to [**StringField**](StringField.md) |  | [optional] 
**ActiveRevisionId** | Pointer to [**GuidNullableField**](GuidNullableField.md) |  | [optional] 

## Methods

### NewPatchChannelCommand

`func NewPatchChannelCommand() *PatchChannelCommand`

NewPatchChannelCommand instantiates a new PatchChannelCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchChannelCommandWithDefaults

`func NewPatchChannelCommandWithDefaults() *PatchChannelCommand`

NewPatchChannelCommandWithDefaults instantiates a new PatchChannelCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *PatchChannelCommand) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PatchChannelCommand) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PatchChannelCommand) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *PatchChannelCommand) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetName

`func (o *PatchChannelCommand) GetName() StringField`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchChannelCommand) GetNameOk() (*StringField, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchChannelCommand) SetName(v StringField)`

SetName sets Name field to given value.

### HasName

`func (o *PatchChannelCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomain

`func (o *PatchChannelCommand) GetDomain() StringField`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PatchChannelCommand) GetDomainOk() (*StringField, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PatchChannelCommand) SetDomain(v StringField)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PatchChannelCommand) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetRevisionSelectionStrategy

`func (o *PatchChannelCommand) GetRevisionSelectionStrategy() ChannelRevisionSelectionStrategyField`

GetRevisionSelectionStrategy returns the RevisionSelectionStrategy field if non-nil, zero value otherwise.

### GetRevisionSelectionStrategyOk

`func (o *PatchChannelCommand) GetRevisionSelectionStrategyOk() (*ChannelRevisionSelectionStrategyField, bool)`

GetRevisionSelectionStrategyOk returns a tuple with the RevisionSelectionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionSelectionStrategy

`func (o *PatchChannelCommand) SetRevisionSelectionStrategy(v ChannelRevisionSelectionStrategyField)`

SetRevisionSelectionStrategy sets RevisionSelectionStrategy field to given value.

### HasRevisionSelectionStrategy

`func (o *PatchChannelCommand) HasRevisionSelectionStrategy() bool`

HasRevisionSelectionStrategy returns a boolean if a field has been set.

### GetRangeRule

`func (o *PatchChannelCommand) GetRangeRule() StringField`

GetRangeRule returns the RangeRule field if non-nil, zero value otherwise.

### GetRangeRuleOk

`func (o *PatchChannelCommand) GetRangeRuleOk() (*StringField, bool)`

GetRangeRuleOk returns a tuple with the RangeRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeRule

`func (o *PatchChannelCommand) SetRangeRule(v StringField)`

SetRangeRule sets RangeRule field to given value.

### HasRangeRule

`func (o *PatchChannelCommand) HasRangeRule() bool`

HasRangeRule returns a boolean if a field has been set.

### GetActiveRevisionId

`func (o *PatchChannelCommand) GetActiveRevisionId() GuidNullableField`

GetActiveRevisionId returns the ActiveRevisionId field if non-nil, zero value otherwise.

### GetActiveRevisionIdOk

`func (o *PatchChannelCommand) GetActiveRevisionIdOk() (*GuidNullableField, bool)`

GetActiveRevisionIdOk returns a tuple with the ActiveRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRevisionId

`func (o *PatchChannelCommand) SetActiveRevisionId(v GuidNullableField)`

SetActiveRevisionId sets ActiveRevisionId field to given value.

### HasActiveRevisionId

`func (o *PatchChannelCommand) HasActiveRevisionId() bool`

HasActiveRevisionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


