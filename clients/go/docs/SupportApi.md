# \SupportApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSupportPost**](SupportApi.md#ApiSupportPost) | **Post** /api/support | 



## ApiSupportPost

> ApiSupportPost(ctx).CreateSupportTicketCommand(createSupportTicketCommand).ApiVersion(apiVersion).Execute()



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
    createSupportTicketCommand := *openapiclient.NewCreateSupportTicketCommand(openapiclient.TicketCategory("TechnicalHelp")) // CreateSupportTicketCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SupportApi.ApiSupportPost(context.Background()).CreateSupportTicketCommand(createSupportTicketCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportApi.ApiSupportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSupportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSupportTicketCommand** | [**CreateSupportTicketCommand**](CreateSupportTicketCommand.md) |  | 
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

