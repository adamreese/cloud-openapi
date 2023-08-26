# CreateAppCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**StorageId** | **string** |  | 
**CreateDefaultDatabase** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateAppCommand

`func NewCreateAppCommand(name string, storageId string, ) *CreateAppCommand`

NewCreateAppCommand instantiates a new CreateAppCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAppCommandWithDefaults

`func NewCreateAppCommandWithDefaults() *CreateAppCommand`

NewCreateAppCommandWithDefaults instantiates a new CreateAppCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAppCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAppCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAppCommand) SetName(v string)`

SetName sets Name field to given value.


### GetStorageId

`func (o *CreateAppCommand) GetStorageId() string`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *CreateAppCommand) GetStorageIdOk() (*string, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *CreateAppCommand) SetStorageId(v string)`

SetStorageId sets StorageId field to given value.


### GetCreateDefaultDatabase

`func (o *CreateAppCommand) GetCreateDefaultDatabase() bool`

GetCreateDefaultDatabase returns the CreateDefaultDatabase field if non-nil, zero value otherwise.

### GetCreateDefaultDatabaseOk

`func (o *CreateAppCommand) GetCreateDefaultDatabaseOk() (*bool, bool)`

GetCreateDefaultDatabaseOk returns a tuple with the CreateDefaultDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDefaultDatabase

`func (o *CreateAppCommand) SetCreateDefaultDatabase(v bool)`

SetCreateDefaultDatabase sets CreateDefaultDatabase field to given value.

### HasCreateDefaultDatabase

`func (o *CreateAppCommand) HasCreateDefaultDatabase() bool`

HasCreateDefaultDatabase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


