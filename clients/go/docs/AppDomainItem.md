# AppDomainItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ValidationStatus** | [**DomainValidationStatus**](DomainValidationStatus.md) |  | 

## Methods

### NewAppDomainItem

`func NewAppDomainItem(name string, validationStatus DomainValidationStatus, ) *AppDomainItem`

NewAppDomainItem instantiates a new AppDomainItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDomainItemWithDefaults

`func NewAppDomainItemWithDefaults() *AppDomainItem`

NewAppDomainItemWithDefaults instantiates a new AppDomainItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppDomainItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDomainItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDomainItem) SetName(v string)`

SetName sets Name field to given value.


### GetValidationStatus

`func (o *AppDomainItem) GetValidationStatus() DomainValidationStatus`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *AppDomainItem) GetValidationStatusOk() (*DomainValidationStatus, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *AppDomainItem) SetValidationStatus(v DomainValidationStatus)`

SetValidationStatus sets ValidationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


