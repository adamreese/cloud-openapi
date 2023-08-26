# AccountPlanRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to [**AccountPlanType**](AccountPlanType.md) |  | [optional] 
**CancelAt** | Pointer to **NullableTime** |  | [optional] 
**CycleStartDate** | Pointer to **NullableTime** |  | [optional] 
**CycleEndDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewAccountPlanRecord

`func NewAccountPlanRecord() *AccountPlanRecord`

NewAccountPlanRecord instantiates a new AccountPlanRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountPlanRecordWithDefaults

`func NewAccountPlanRecordWithDefaults() *AccountPlanRecord`

NewAccountPlanRecordWithDefaults instantiates a new AccountPlanRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *AccountPlanRecord) GetAccountType() AccountPlanType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountPlanRecord) GetAccountTypeOk() (*AccountPlanType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountPlanRecord) SetAccountType(v AccountPlanType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountPlanRecord) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetCancelAt

`func (o *AccountPlanRecord) GetCancelAt() time.Time`

GetCancelAt returns the CancelAt field if non-nil, zero value otherwise.

### GetCancelAtOk

`func (o *AccountPlanRecord) GetCancelAtOk() (*time.Time, bool)`

GetCancelAtOk returns a tuple with the CancelAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAt

`func (o *AccountPlanRecord) SetCancelAt(v time.Time)`

SetCancelAt sets CancelAt field to given value.

### HasCancelAt

`func (o *AccountPlanRecord) HasCancelAt() bool`

HasCancelAt returns a boolean if a field has been set.

### SetCancelAtNil

`func (o *AccountPlanRecord) SetCancelAtNil(b bool)`

 SetCancelAtNil sets the value for CancelAt to be an explicit nil

### UnsetCancelAt
`func (o *AccountPlanRecord) UnsetCancelAt()`

UnsetCancelAt ensures that no value is present for CancelAt, not even an explicit nil
### GetCycleStartDate

`func (o *AccountPlanRecord) GetCycleStartDate() time.Time`

GetCycleStartDate returns the CycleStartDate field if non-nil, zero value otherwise.

### GetCycleStartDateOk

`func (o *AccountPlanRecord) GetCycleStartDateOk() (*time.Time, bool)`

GetCycleStartDateOk returns a tuple with the CycleStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleStartDate

`func (o *AccountPlanRecord) SetCycleStartDate(v time.Time)`

SetCycleStartDate sets CycleStartDate field to given value.

### HasCycleStartDate

`func (o *AccountPlanRecord) HasCycleStartDate() bool`

HasCycleStartDate returns a boolean if a field has been set.

### SetCycleStartDateNil

`func (o *AccountPlanRecord) SetCycleStartDateNil(b bool)`

 SetCycleStartDateNil sets the value for CycleStartDate to be an explicit nil

### UnsetCycleStartDate
`func (o *AccountPlanRecord) UnsetCycleStartDate()`

UnsetCycleStartDate ensures that no value is present for CycleStartDate, not even an explicit nil
### GetCycleEndDate

`func (o *AccountPlanRecord) GetCycleEndDate() time.Time`

GetCycleEndDate returns the CycleEndDate field if non-nil, zero value otherwise.

### GetCycleEndDateOk

`func (o *AccountPlanRecord) GetCycleEndDateOk() (*time.Time, bool)`

GetCycleEndDateOk returns a tuple with the CycleEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleEndDate

`func (o *AccountPlanRecord) SetCycleEndDate(v time.Time)`

SetCycleEndDate sets CycleEndDate field to given value.

### HasCycleEndDate

`func (o *AccountPlanRecord) HasCycleEndDate() bool`

HasCycleEndDate returns a boolean if a field has been set.

### SetCycleEndDateNil

`func (o *AccountPlanRecord) SetCycleEndDateNil(b bool)`

 SetCycleEndDateNil sets the value for CycleEndDate to be an explicit nil

### UnsetCycleEndDate
`func (o *AccountPlanRecord) UnsetCycleEndDate()`

UnsetCycleEndDate ensures that no value is present for CycleEndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


