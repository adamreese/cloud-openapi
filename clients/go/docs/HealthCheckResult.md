# HealthCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ApiHealthStatus**](ApiHealthStatus.md) |  | [optional] 

## Methods

### NewHealthCheckResult

`func NewHealthCheckResult() *HealthCheckResult`

NewHealthCheckResult instantiates a new HealthCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckResultWithDefaults

`func NewHealthCheckResultWithDefaults() *HealthCheckResult`

NewHealthCheckResultWithDefaults instantiates a new HealthCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthCheckResult) GetStatus() ApiHealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckResult) GetStatusOk() (*ApiHealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckResult) SetStatus(v ApiHealthStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthCheckResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


