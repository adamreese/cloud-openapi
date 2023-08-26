# \VariablePairsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiVariablePairsDelete**](VariablePairsApi.md#ApiVariablePairsDelete) | **Delete** /api/variable-pairs | 
[**ApiVariablePairsGet**](VariablePairsApi.md#ApiVariablePairsGet) | **Get** /api/variable-pairs | 
[**ApiVariablePairsPost**](VariablePairsApi.md#ApiVariablePairsPost) | **Post** /api/variable-pairs | 



## ApiVariablePairsDelete

> ApiVariablePairsDelete(ctx).DeleteVariablePairCommand(deleteVariablePairCommand).ApiVersion(apiVersion).Execute()



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
    deleteVariablePairCommand := *openapiclient.NewDeleteVariablePairCommand("AppId_example", "Variable_example") // DeleteVariablePairCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VariablePairsApi.ApiVariablePairsDelete(context.Background()).DeleteVariablePairCommand(deleteVariablePairCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablePairsApi.ApiVariablePairsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiVariablePairsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVariablePairCommand** | [**DeleteVariablePairCommand**](DeleteVariablePairCommand.md) |  | 
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


## ApiVariablePairsGet

> VariablesList ApiVariablePairsGet(ctx).GetVariablesQuery(getVariablesQuery).ApiVersion(apiVersion).Execute()



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
    getVariablesQuery := *openapiclient.NewGetVariablesQuery("AppId_example") // GetVariablesQuery | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablePairsApi.ApiVariablePairsGet(context.Background()).GetVariablesQuery(getVariablesQuery).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablePairsApi.ApiVariablePairsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiVariablePairsGet`: VariablesList
    fmt.Fprintf(os.Stdout, "Response from `VariablePairsApi.ApiVariablePairsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiVariablePairsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getVariablesQuery** | [**GetVariablesQuery**](GetVariablesQuery.md) |  | 
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**VariablesList**](VariablesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiVariablePairsPost

> ApiVariablePairsPost(ctx).CreateVariablePairCommand(createVariablePairCommand).ApiVersion(apiVersion).Execute()



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
    createVariablePairCommand := *openapiclient.NewCreateVariablePairCommand("AppId_example", "Variable_example", "Value_example") // CreateVariablePairCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VariablePairsApi.ApiVariablePairsPost(context.Background()).CreateVariablePairCommand(createVariablePairCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablePairsApi.ApiVariablePairsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiVariablePairsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVariablePairCommand** | [**CreateVariablePairCommand**](CreateVariablePairCommand.md) |  | 
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

