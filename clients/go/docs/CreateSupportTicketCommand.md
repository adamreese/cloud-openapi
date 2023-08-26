# CreateSupportTicketCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**TicketCategory**](TicketCategory.md) |  | 
**Details** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSupportTicketCommand

`func NewCreateSupportTicketCommand(category TicketCategory, ) *CreateSupportTicketCommand`

NewCreateSupportTicketCommand instantiates a new CreateSupportTicketCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSupportTicketCommandWithDefaults

`func NewCreateSupportTicketCommandWithDefaults() *CreateSupportTicketCommand`

NewCreateSupportTicketCommandWithDefaults instantiates a new CreateSupportTicketCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CreateSupportTicketCommand) GetCategory() TicketCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateSupportTicketCommand) GetCategoryOk() (*TicketCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateSupportTicketCommand) SetCategory(v TicketCategory)`

SetCategory sets Category field to given value.


### GetDetails

`func (o *CreateSupportTicketCommand) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CreateSupportTicketCommand) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CreateSupportTicketCommand) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CreateSupportTicketCommand) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


