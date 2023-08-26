# PersonalAccessTokenItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPersonalAccessTokenItem

`func NewPersonalAccessTokenItem() *PersonalAccessTokenItem`

NewPersonalAccessTokenItem instantiates a new PersonalAccessTokenItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalAccessTokenItemWithDefaults

`func NewPersonalAccessTokenItemWithDefaults() *PersonalAccessTokenItem`

NewPersonalAccessTokenItemWithDefaults instantiates a new PersonalAccessTokenItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonalAccessTokenItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonalAccessTokenItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonalAccessTokenItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PersonalAccessTokenItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PersonalAccessTokenItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersonalAccessTokenItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersonalAccessTokenItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PersonalAccessTokenItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PersonalAccessTokenItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PersonalAccessTokenItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PersonalAccessTokenItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PersonalAccessTokenItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


