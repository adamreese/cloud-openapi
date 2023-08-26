# AppChannelListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ActiveRevisionNumber** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Created** | **time.Time** |  | 
**HealthStatus** | Pointer to [**ApiHealthStatus**](ApiHealthStatus.md) |  | [optional] 

## Methods

### NewAppChannelListItem

`func NewAppChannelListItem(id string, name string, created time.Time, ) *AppChannelListItem`

NewAppChannelListItem instantiates a new AppChannelListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppChannelListItemWithDefaults

`func NewAppChannelListItemWithDefaults() *AppChannelListItem`

NewAppChannelListItemWithDefaults instantiates a new AppChannelListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppChannelListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppChannelListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppChannelListItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AppChannelListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppChannelListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppChannelListItem) SetName(v string)`

SetName sets Name field to given value.


### GetActiveRevisionNumber

`func (o *AppChannelListItem) GetActiveRevisionNumber() string`

GetActiveRevisionNumber returns the ActiveRevisionNumber field if non-nil, zero value otherwise.

### GetActiveRevisionNumberOk

`func (o *AppChannelListItem) GetActiveRevisionNumberOk() (*string, bool)`

GetActiveRevisionNumberOk returns a tuple with the ActiveRevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRevisionNumber

`func (o *AppChannelListItem) SetActiveRevisionNumber(v string)`

SetActiveRevisionNumber sets ActiveRevisionNumber field to given value.

### HasActiveRevisionNumber

`func (o *AppChannelListItem) HasActiveRevisionNumber() bool`

HasActiveRevisionNumber returns a boolean if a field has been set.

### GetDomain

`func (o *AppChannelListItem) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AppChannelListItem) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AppChannelListItem) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AppChannelListItem) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetCreated

`func (o *AppChannelListItem) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppChannelListItem) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppChannelListItem) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetHealthStatus

`func (o *AppChannelListItem) GetHealthStatus() ApiHealthStatus`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *AppChannelListItem) GetHealthStatusOk() (*ApiHealthStatus, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *AppChannelListItem) SetHealthStatus(v ApiHealthStatus)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *AppChannelListItem) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


