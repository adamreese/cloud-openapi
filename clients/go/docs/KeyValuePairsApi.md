# \KeyValuePairsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeyValuePairsPost**](KeyValuePairsApi.md#ApiKeyValuePairsPost) | **Post** /api/key-value-pairs | 



## ApiKeyValuePairsPost

> ApiKeyValuePairsPost(ctx).CreateKeyValuePairCommand(createKeyValuePairCommand).ApiVersion(apiVersion).Execute()



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
    createKeyValuePairCommand := *openapiclient.NewCreateKeyValuePairCommand("AppId_example", "StoreName_example", "Key_example", "Value_example") // CreateKeyValuePairCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyValuePairsApi.ApiKeyValuePairsPost(context.Background()).CreateKeyValuePairCommand(createKeyValuePairCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyValuePairsApi.ApiKeyValuePairsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeyValuePairsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKeyValuePairCommand** | [**CreateKeyValuePairCommand**](CreateKeyValuePairCommand.md) |  | 
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

