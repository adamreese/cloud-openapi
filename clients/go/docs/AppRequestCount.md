# AppRequestCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **int64** |  | 
**End** | **int64** |  | 
**Interval** | **int32** |  | 
**Total** | **int64** |  | 
**AveragePerSecond** | **float64** |  | 
**Points** | [**[]AppRequestPoint**](AppRequestPoint.md) |  | 

## Methods

### NewAppRequestCount

`func NewAppRequestCount(start int64, end int64, interval int32, total int64, averagePerSecond float64, points []AppRequestPoint, ) *AppRequestCount`

NewAppRequestCount instantiates a new AppRequestCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRequestCountWithDefaults

`func NewAppRequestCountWithDefaults() *AppRequestCount`

NewAppRequestCountWithDefaults instantiates a new AppRequestCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *AppRequestCount) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AppRequestCount) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AppRequestCount) SetStart(v int64)`

SetStart sets Start field to given value.


### GetEnd

`func (o *AppRequestCount) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AppRequestCount) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AppRequestCount) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetInterval

`func (o *AppRequestCount) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AppRequestCount) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AppRequestCount) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetTotal

`func (o *AppRequestCount) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AppRequestCount) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AppRequestCount) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetAveragePerSecond

`func (o *AppRequestCount) GetAveragePerSecond() float64`

GetAveragePerSecond returns the AveragePerSecond field if non-nil, zero value otherwise.

### GetAveragePerSecondOk

`func (o *AppRequestCount) GetAveragePerSecondOk() (*float64, bool)`

GetAveragePerSecondOk returns a tuple with the AveragePerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePerSecond

`func (o *AppRequestCount) SetAveragePerSecond(v float64)`

SetAveragePerSecond sets AveragePerSecond field to given value.


### GetPoints

`func (o *AppRequestCount) GetPoints() []AppRequestPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *AppRequestCount) GetPointsOk() (*[]AppRequestPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *AppRequestCount) SetPoints(v []AppRequestPoint)`

SetPoints sets Points field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


