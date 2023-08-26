# \DeviceCodesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDeviceCodesActivatePost**](DeviceCodesApi.md#ApiDeviceCodesActivatePost) | **Post** /api/device-codes/activate | 
[**ApiDeviceCodesPost**](DeviceCodesApi.md#ApiDeviceCodesPost) | **Post** /api/device-codes | 
[**ApiDeviceCodesUserCodeGet**](DeviceCodesApi.md#ApiDeviceCodesUserCodeGet) | **Get** /api/device-codes/{userCode} | 



## ApiDeviceCodesActivatePost

> ApiDeviceCodesActivatePost(ctx).ActivateDeviceCodeCommand(activateDeviceCodeCommand).ApiVersion(apiVersion).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fermyon/cloud-openapi"
)

func main() {
    activateDeviceCodeCommand := *openapiclient.NewActivateDeviceCodeCommand("UserCode_example") // ActivateDeviceCodeCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeviceCodesApi.ApiDeviceCodesActivatePost(context.Background()).ActivateDeviceCodeCommand(activateDeviceCodeCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceCodesApi.ApiDeviceCodesActivatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDeviceCodesActivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activateDeviceCodeCommand** | [**ActivateDeviceCodeCommand**](ActivateDeviceCodeCommand.md) |  | 
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeviceCodesPost

> DeviceCodeItem ApiDeviceCodesPost(ctx).CreateDeviceCodeCommand(createDeviceCodeCommand).ApiVersion(apiVersion).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fermyon/cloud-openapi"
)

func main() {
    createDeviceCodeCommand := *openapiclient.NewCreateDeviceCodeCommand("ClientId_example") // CreateDeviceCodeCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceCodesApi.ApiDeviceCodesPost(context.Background()).CreateDeviceCodeCommand(createDeviceCodeCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceCodesApi.ApiDeviceCodesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeviceCodesPost`: DeviceCodeItem
    fmt.Fprintf(os.Stdout, "Response from `DeviceCodesApi.ApiDeviceCodesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDeviceCodesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDeviceCodeCommand** | [**CreateDeviceCodeCommand**](CreateDeviceCodeCommand.md) |  | 
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**DeviceCodeItem**](DeviceCodeItem.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDeviceCodesUserCodeGet

> DeviceCodeDetails ApiDeviceCodesUserCodeGet(ctx, userCode).ApiVersion(apiVersion).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fermyon/cloud-openapi"
)

func main() {
    userCode := "userCode_example" // string | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceCodesApi.ApiDeviceCodesUserCodeGet(context.Background(), userCode).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceCodesApi.ApiDeviceCodesUserCodeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDeviceCodesUserCodeGet`: DeviceCodeDetails
    fmt.Fprintf(os.Stdout, "Response from `DeviceCodesApi.ApiDeviceCodesUserCodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDeviceCodesUserCodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**DeviceCodeDetails**](DeviceCodeDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

