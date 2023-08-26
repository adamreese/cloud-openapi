# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountPlanType** | Pointer to [**AccountPlanType**](AccountPlanType.md) |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewPlan

`func NewPlan() *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountPlanType

`func (o *Plan) GetAccountPlanType() AccountPlanType`

GetAccountPlanType returns the AccountPlanType field if non-nil, zero value otherwise.

### GetAccountPlanTypeOk

`func (o *Plan) GetAccountPlanTypeOk() (*AccountPlanType, bool)`

GetAccountPlanTypeOk returns a tuple with the AccountPlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPlanType

`func (o *Plan) SetAccountPlanType(v AccountPlanType)`

SetAccountPlanType sets AccountPlanType field to given value.

### HasAccountPlanType

`func (o *Plan) HasAccountPlanType() bool`

HasAccountPlanType returns a boolean if a field has been set.

### GetPrice

`func (o *Plan) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Plan) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Plan) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Plan) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *Plan) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *Plan) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


