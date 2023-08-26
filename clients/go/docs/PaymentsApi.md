# \PaymentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPaymentsCustomerPortalGet**](PaymentsApi.md#ApiPaymentsCustomerPortalGet) | **Get** /api/payments/customer-portal | 
[**ApiPaymentsPlansGet**](PaymentsApi.md#ApiPaymentsPlansGet) | **Get** /api/payments/plans | 
[**ApiPaymentsSetupCheckoutPost**](PaymentsApi.md#ApiPaymentsSetupCheckoutPost) | **Post** /api/payments/setup-checkout | 



## ApiPaymentsCustomerPortalGet

> PaymentIntegrationUrl ApiPaymentsCustomerPortalGet(ctx).ApiVersion(apiVersion).Execute()



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
    resp, r, err := apiClient.PaymentsApi.ApiPaymentsCustomerPortalGet(context.Background()).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.ApiPaymentsCustomerPortalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPaymentsCustomerPortalGet`: PaymentIntegrationUrl
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.ApiPaymentsCustomerPortalGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPaymentsCustomerPortalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**PaymentIntegrationUrl**](PaymentIntegrationUrl.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPaymentsPlansGet

> []Plan ApiPaymentsPlansGet(ctx).ApiVersion(apiVersion).Execute()



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
    resp, r, err := apiClient.PaymentsApi.ApiPaymentsPlansGet(context.Background()).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.ApiPaymentsPlansGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPaymentsPlansGet`: []Plan
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.ApiPaymentsPlansGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPaymentsPlansGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**[]Plan**](Plan.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPaymentsSetupCheckoutPost

> PaymentIntegrationUrl ApiPaymentsSetupCheckoutPost(ctx).ApiVersion(apiVersion).Execute()



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
    resp, r, err := apiClient.PaymentsApi.ApiPaymentsSetupCheckoutPost(context.Background()).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.ApiPaymentsSetupCheckoutPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPaymentsSetupCheckoutPost`: PaymentIntegrationUrl
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.ApiPaymentsSetupCheckoutPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPaymentsSetupCheckoutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**PaymentIntegrationUrl**](PaymentIntegrationUrl.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

