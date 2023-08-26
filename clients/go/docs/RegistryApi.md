# \RegistryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRegistryIBindleNameGet**](RegistryApi.md#ApiRegistryIBindleNameGet) | **Get** /api/registry/_i/{bindleName} | 
[**ApiRegistryIBindleNamePost**](RegistryApi.md#ApiRegistryIBindleNamePost) | **Post** /api/registry/_i/{bindleName} | 
[**ApiRegistryIPost**](RegistryApi.md#ApiRegistryIPost) | **Post** /api/registry/_i | 



## ApiRegistryIBindleNameGet

> ApiRegistryIBindleNameGet(ctx, bindleName).ApiVersion(apiVersion).Execute()



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
    bindleName := "bindleName_example" // string | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RegistryApi.ApiRegistryIBindleNameGet(context.Background(), bindleName).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ApiRegistryIBindleNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bindleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRegistryIBindleNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRegistryIBindleNamePost

> ApiRegistryIBindleNamePost(ctx, bindleName).ApiVersion(apiVersion).Execute()



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
    bindleName := "bindleName_example" // string | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RegistryApi.ApiRegistryIBindleNamePost(context.Background(), bindleName).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ApiRegistryIBindleNamePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bindleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRegistryIBindleNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRegistryIPost

> ApiRegistryIPost(ctx).ApiVersion(apiVersion).Execute()



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
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RegistryApi.ApiRegistryIPost(context.Background()).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ApiRegistryIPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRegistryIPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

