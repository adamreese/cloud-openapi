# DomainItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ValidationStatus** | **string** |  | 
**ValidatedAt** | Pointer to **NullableTime** |  | [optional] 
**DnsRecords** | [**[]DnsRecord**](DnsRecord.md) |  | 

## Methods

### NewDomainItem

`func NewDomainItem(name string, validationStatus string, dnsRecords []DnsRecord, ) *DomainItem`

NewDomainItem instantiates a new DomainItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainItemWithDefaults

`func NewDomainItemWithDefaults() *DomainItem`

NewDomainItemWithDefaults instantiates a new DomainItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DomainItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainItem) SetName(v string)`

SetName sets Name field to given value.


### GetValidationStatus

`func (o *DomainItem) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *DomainItem) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *DomainItem) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.


### GetValidatedAt

`func (o *DomainItem) GetValidatedAt() time.Time`

GetValidatedAt returns the ValidatedAt field if non-nil, zero value otherwise.

### GetValidatedAtOk

`func (o *DomainItem) GetValidatedAtOk() (*time.Time, bool)`

GetValidatedAtOk returns a tuple with the ValidatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatedAt

`func (o *DomainItem) SetValidatedAt(v time.Time)`

SetValidatedAt sets ValidatedAt field to given value.

### HasValidatedAt

`func (o *DomainItem) HasValidatedAt() bool`

HasValidatedAt returns a boolean if a field has been set.

### SetValidatedAtNil

`func (o *DomainItem) SetValidatedAtNil(b bool)`

 SetValidatedAtNil sets the value for ValidatedAt to be an explicit nil

### UnsetValidatedAt
`func (o *DomainItem) UnsetValidatedAt()`

UnsetValidatedAt ensures that no value is present for ValidatedAt, not even an explicit nil
### GetDnsRecords

`func (o *DomainItem) GetDnsRecords() []DnsRecord`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *DomainItem) GetDnsRecordsOk() (*[]DnsRecord, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *DomainItem) SetDnsRecords(v []DnsRecord)`

SetDnsRecords sets DnsRecords field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


