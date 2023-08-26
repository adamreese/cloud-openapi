# CreateChannelCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** |  | 
**Name** | **string** |  | 
**RevisionSelectionStrategy** | [**ChannelRevisionSelectionStrategy**](ChannelRevisionSelectionStrategy.md) |  | 
**RangeRule** | Pointer to **NullableString** |  | [optional] 
**ActiveRevisionId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateChannelCommand

`func NewCreateChannelCommand(appId string, name string, revisionSelectionStrategy ChannelRevisionSelectionStrategy, ) *CreateChannelCommand`

NewCreateChannelCommand instantiates a new CreateChannelCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateChannelCommandWithDefaults

`func NewCreateChannelCommandWithDefaults() *CreateChannelCommand`

NewCreateChannelCommandWithDefaults instantiates a new CreateChannelCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *CreateChannelCommand) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateChannelCommand) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateChannelCommand) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetName

`func (o *CreateChannelCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateChannelCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateChannelCommand) SetName(v string)`

SetName sets Name field to given value.


### GetRevisionSelectionStrategy

`func (o *CreateChannelCommand) GetRevisionSelectionStrategy() ChannelRevisionSelectionStrategy`

GetRevisionSelectionStrategy returns the RevisionSelectionStrategy field if non-nil, zero value otherwise.

### GetRevisionSelectionStrategyOk

`func (o *CreateChannelCommand) GetRevisionSelectionStrategyOk() (*ChannelRevisionSelectionStrategy, bool)`

GetRevisionSelectionStrategyOk returns a tuple with the RevisionSelectionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionSelectionStrategy

`func (o *CreateChannelCommand) SetRevisionSelectionStrategy(v ChannelRevisionSelectionStrategy)`

SetRevisionSelectionStrategy sets RevisionSelectionStrategy field to given value.


### GetRangeRule

`func (o *CreateChannelCommand) GetRangeRule() string`

GetRangeRule returns the RangeRule field if non-nil, zero value otherwise.

### GetRangeRuleOk

`func (o *CreateChannelCommand) GetRangeRuleOk() (*string, bool)`

GetRangeRuleOk returns a tuple with the RangeRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeRule

`func (o *CreateChannelCommand) SetRangeRule(v string)`

SetRangeRule sets RangeRule field to given value.

### HasRangeRule

`func (o *CreateChannelCommand) HasRangeRule() bool`

HasRangeRule returns a boolean if a field has been set.

### SetRangeRuleNil

`func (o *CreateChannelCommand) SetRangeRuleNil(b bool)`

 SetRangeRuleNil sets the value for RangeRule to be an explicit nil

### UnsetRangeRule
`func (o *CreateChannelCommand) UnsetRangeRule()`

UnsetRangeRule ensures that no value is present for RangeRule, not even an explicit nil
### GetActiveRevisionId

`func (o *CreateChannelCommand) GetActiveRevisionId() string`

GetActiveRevisionId returns the ActiveRevisionId field if non-nil, zero value otherwise.

### GetActiveRevisionIdOk

`func (o *CreateChannelCommand) GetActiveRevisionIdOk() (*string, bool)`

GetActiveRevisionIdOk returns a tuple with the ActiveRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRevisionId

`func (o *CreateChannelCommand) SetActiveRevisionId(v string)`

SetActiveRevisionId sets ActiveRevisionId field to given value.

### HasActiveRevisionId

`func (o *CreateChannelCommand) HasActiveRevisionId() bool`

HasActiveRevisionId returns a boolean if a field has been set.

### SetActiveRevisionIdNil

`func (o *CreateChannelCommand) SetActiveRevisionIdNil(b bool)`

 SetActiveRevisionIdNil sets the value for ActiveRevisionId to be an explicit nil

### UnsetActiveRevisionId
`func (o *CreateChannelCommand) UnsetActiveRevisionId()`

UnsetActiveRevisionId ensures that no value is present for ActiveRevisionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


