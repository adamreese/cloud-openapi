# AccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | Pointer to [**AccountPlanRecord**](AccountPlanRecord.md) |  | [optional] 
**IsMarketingEmailOn** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAccountDetails

`func NewAccountDetails() *AccountDetails`

NewAccountDetails instantiates a new AccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsWithDefaults

`func NewAccountDetailsWithDefaults() *AccountDetails`

NewAccountDetailsWithDefaults instantiates a new AccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *AccountDetails) GetPlan() AccountPlanRecord`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AccountDetails) GetPlanOk() (*AccountPlanRecord, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AccountDetails) SetPlan(v AccountPlanRecord)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AccountDetails) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetIsMarketingEmailOn

`func (o *AccountDetails) GetIsMarketingEmailOn() bool`

GetIsMarketingEmailOn returns the IsMarketingEmailOn field if non-nil, zero value otherwise.

### GetIsMarketingEmailOnOk

`func (o *AccountDetails) GetIsMarketingEmailOnOk() (*bool, bool)`

GetIsMarketingEmailOnOk returns a tuple with the IsMarketingEmailOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketingEmailOn

`func (o *AccountDetails) SetIsMarketingEmailOn(v bool)`

SetIsMarketingEmailOn sets IsMarketingEmailOn field to given value.

### HasIsMarketingEmailOn

`func (o *AccountDetails) HasIsMarketingEmailOn() bool`

HasIsMarketingEmailOn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *AccountDetails) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AccountDetails) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


