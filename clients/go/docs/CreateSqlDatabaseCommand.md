# CreateSqlDatabaseCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AppId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateSqlDatabaseCommand

`func NewCreateSqlDatabaseCommand(name string, ) *CreateSqlDatabaseCommand`

NewCreateSqlDatabaseCommand instantiates a new CreateSqlDatabaseCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSqlDatabaseCommandWithDefaults

`func NewCreateSqlDatabaseCommandWithDefaults() *CreateSqlDatabaseCommand`

NewCreateSqlDatabaseCommandWithDefaults instantiates a new CreateSqlDatabaseCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSqlDatabaseCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSqlDatabaseCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSqlDatabaseCommand) SetName(v string)`

SetName sets Name field to given value.


### GetAppId

`func (o *CreateSqlDatabaseCommand) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *CreateSqlDatabaseCommand) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *CreateSqlDatabaseCommand) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *CreateSqlDatabaseCommand) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *CreateSqlDatabaseCommand) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *CreateSqlDatabaseCommand) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


