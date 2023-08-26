# PersonalAccessTokenItemPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]PersonalAccessTokenItem**](PersonalAccessTokenItem.md) |  | 
**TotalItems** | **int32** |  | 
**PageIndex** | **int32** |  | 
**PageSize** | **int32** |  | 
**IsLastPage** | **bool** |  | [readonly] 

## Methods

### NewPersonalAccessTokenItemPage

`func NewPersonalAccessTokenItemPage(items []PersonalAccessTokenItem, totalItems int32, pageIndex int32, pageSize int32, isLastPage bool, ) *PersonalAccessTokenItemPage`

NewPersonalAccessTokenItemPage instantiates a new PersonalAccessTokenItemPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalAccessTokenItemPageWithDefaults

`func NewPersonalAccessTokenItemPageWithDefaults() *PersonalAccessTokenItemPage`

NewPersonalAccessTokenItemPageWithDefaults instantiates a new PersonalAccessTokenItemPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PersonalAccessTokenItemPage) GetItems() []PersonalAccessTokenItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PersonalAccessTokenItemPage) GetItemsOk() (*[]PersonalAccessTokenItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PersonalAccessTokenItemPage) SetItems(v []PersonalAccessTokenItem)`

SetItems sets Items field to given value.


### GetTotalItems

`func (o *PersonalAccessTokenItemPage) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *PersonalAccessTokenItemPage) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *PersonalAccessTokenItemPage) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.


### GetPageIndex

`func (o *PersonalAccessTokenItemPage) GetPageIndex() int32`

GetPageIndex returns the PageIndex field if non-nil, zero value otherwise.

### GetPageIndexOk

`func (o *PersonalAccessTokenItemPage) GetPageIndexOk() (*int32, bool)`

GetPageIndexOk returns a tuple with the PageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageIndex

`func (o *PersonalAccessTokenItemPage) SetPageIndex(v int32)`

SetPageIndex sets PageIndex field to given value.


### GetPageSize

`func (o *PersonalAccessTokenItemPage) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PersonalAccessTokenItemPage) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PersonalAccessTokenItemPage) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetIsLastPage

`func (o *PersonalAccessTokenItemPage) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *PersonalAccessTokenItemPage) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *PersonalAccessTokenItemPage) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


