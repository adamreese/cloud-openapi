# CreateTokenCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderCode** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**AccountProvider**](AccountProvider.md) |  | [optional] 
**IsMarketingEmailOn** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateTokenCommand

`func NewCreateTokenCommand() *CreateTokenCommand`

NewCreateTokenCommand instantiates a new CreateTokenCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenCommandWithDefaults

`func NewCreateTokenCommandWithDefaults() *CreateTokenCommand`

NewCreateTokenCommandWithDefaults instantiates a new CreateTokenCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderCode

`func (o *CreateTokenCommand) GetProviderCode() string`

GetProviderCode returns the ProviderCode field if non-nil, zero value otherwise.

### GetProviderCodeOk

`func (o *CreateTokenCommand) GetProviderCodeOk() (*string, bool)`

GetProviderCodeOk returns a tuple with the ProviderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCode

`func (o *CreateTokenCommand) SetProviderCode(v string)`

SetProviderCode sets ProviderCode field to given value.

### HasProviderCode

`func (o *CreateTokenCommand) HasProviderCode() bool`

HasProviderCode returns a boolean if a field has been set.

### GetClientId

`func (o *CreateTokenCommand) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateTokenCommand) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateTokenCommand) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateTokenCommand) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetProvider

`func (o *CreateTokenCommand) GetProvider() AccountProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateTokenCommand) GetProviderOk() (*AccountProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateTokenCommand) SetProvider(v AccountProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CreateTokenCommand) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetIsMarketingEmailOn

`func (o *CreateTokenCommand) GetIsMarketingEmailOn() bool`

GetIsMarketingEmailOn returns the IsMarketingEmailOn field if non-nil, zero value otherwise.

### GetIsMarketingEmailOnOk

`func (o *CreateTokenCommand) GetIsMarketingEmailOnOk() (*bool, bool)`

GetIsMarketingEmailOnOk returns a tuple with the IsMarketingEmailOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketingEmailOn

`func (o *CreateTokenCommand) SetIsMarketingEmailOn(v bool)`

SetIsMarketingEmailOn sets IsMarketingEmailOn field to given value.

### HasIsMarketingEmailOn

`func (o *CreateTokenCommand) HasIsMarketingEmailOn() bool`

HasIsMarketingEmailOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


