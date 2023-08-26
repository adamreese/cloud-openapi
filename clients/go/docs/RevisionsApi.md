# \RevisionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRevisionsGet**](RevisionsApi.md#ApiRevisionsGet) | **Get** /api/revisions | 
[**ApiRevisionsPost**](RevisionsApi.md#ApiRevisionsPost) | **Post** /api/revisions | 



## ApiRevisionsGet

> RevisionItemPage ApiRevisionsGet(ctx).PageIndex(pageIndex).PageSize(pageSize).SearchText(searchText).ApiVersion(apiVersion).Execute()



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
    pageIndex := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 50)
    searchText := "searchText_example" // string |  (optional) (default to "")
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionsApi.ApiRevisionsGet(context.Background()).PageIndex(pageIndex).PageSize(pageSize).SearchText(searchText).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionsApi.ApiRevisionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRevisionsGet`: RevisionItemPage
    fmt.Fprintf(os.Stdout, "Response from `RevisionsApi.ApiRevisionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRevisionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageIndex** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 50]
 **searchText** | **string** |  | [default to &quot;&quot;]
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**RevisionItemPage**](RevisionItemPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRevisionsPost

> ApiRevisionsPost(ctx).RegisterRevisionCommand(registerRevisionCommand).ApiVersion(apiVersion).Execute()



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
    registerRevisionCommand := *openapiclient.NewRegisterRevisionCommand("AppStorageId_example", "RevisionNumber_example") // RegisterRevisionCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RevisionsApi.ApiRevisionsPost(context.Background()).RegisterRevisionCommand(registerRevisionCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionsApi.ApiRevisionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRevisionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerRevisionCommand** | [**RegisterRevisionCommand**](RegisterRevisionCommand.md) |  | 
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

