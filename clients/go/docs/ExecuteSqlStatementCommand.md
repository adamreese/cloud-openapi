# ExecuteSqlStatementCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | **string** |  | 
**Default** | **bool** |  | 
**Statement** | **string** |  | 

## Methods

### NewExecuteSqlStatementCommand

`func NewExecuteSqlStatementCommand(database string, default_ bool, statement string, ) *ExecuteSqlStatementCommand`

NewExecuteSqlStatementCommand instantiates a new ExecuteSqlStatementCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteSqlStatementCommandWithDefaults

`func NewExecuteSqlStatementCommandWithDefaults() *ExecuteSqlStatementCommand`

NewExecuteSqlStatementCommandWithDefaults instantiates a new ExecuteSqlStatementCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *ExecuteSqlStatementCommand) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ExecuteSqlStatementCommand) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ExecuteSqlStatementCommand) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetDefault

`func (o *ExecuteSqlStatementCommand) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ExecuteSqlStatementCommand) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ExecuteSqlStatementCommand) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetStatement

`func (o *ExecuteSqlStatementCommand) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *ExecuteSqlStatementCommand) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *ExecuteSqlStatementCommand) SetStatement(v string)`

SetStatement sets Statement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


