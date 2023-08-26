# \PersonalAccessTokensApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPersonalAccessTokensGet**](PersonalAccessTokensApi.md#ApiPersonalAccessTokensGet) | **Get** /api/personal-access-tokens | 
[**ApiPersonalAccessTokensIdDelete**](PersonalAccessTokensApi.md#ApiPersonalAccessTokensIdDelete) | **Delete** /api/personal-access-tokens/{id} | 
[**ApiPersonalAccessTokensPost**](PersonalAccessTokensApi.md#ApiPersonalAccessTokensPost) | **Post** /api/personal-access-tokens | 



## ApiPersonalAccessTokensGet

> PersonalAccessTokenItemPage ApiPersonalAccessTokensGet(ctx).SearchText(searchText).PageIndex(pageIndex).PageSize(pageSize).SortBy(sortBy).IsSortedAscending(isSortedAscending).ApiVersion(apiVersion).Execute()



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
    searchText := "searchText_example" // string |  (optional) (default to "")
    pageIndex := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 50)
    sortBy := "sortBy_example" // string |  (optional) (default to "CreatedAt")
    isSortedAscending := true // bool |  (optional) (default to false)
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalAccessTokensApi.ApiPersonalAccessTokensGet(context.Background()).SearchText(searchText).PageIndex(pageIndex).PageSize(pageSize).SortBy(sortBy).IsSortedAscending(isSortedAscending).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensApi.ApiPersonalAccessTokensGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPersonalAccessTokensGet`: PersonalAccessTokenItemPage
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokensApi.ApiPersonalAccessTokensGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPersonalAccessTokensGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchText** | **string** |  | [default to &quot;&quot;]
 **pageIndex** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 50]
 **sortBy** | **string** |  | [default to &quot;CreatedAt&quot;]
 **isSortedAscending** | **bool** |  | [default to false]
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**PersonalAccessTokenItemPage**](PersonalAccessTokenItemPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPersonalAccessTokensIdDelete

> ApiPersonalAccessTokensIdDelete(ctx, id).ApiVersion(apiVersion).Execute()



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PersonalAccessTokensApi.ApiPersonalAccessTokensIdDelete(context.Background(), id).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensApi.ApiPersonalAccessTokensIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPersonalAccessTokensIdDeleteRequest struct via the builder pattern


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


## ApiPersonalAccessTokensPost

> PersonalAccessTokenValue ApiPersonalAccessTokensPost(ctx).CreatePersonalAccessTokenCommand(createPersonalAccessTokenCommand).ApiVersion(apiVersion).Execute()



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
    createPersonalAccessTokenCommand := *openapiclient.NewCreatePersonalAccessTokenCommand("Name_example") // CreatePersonalAccessTokenCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalAccessTokensApi.ApiPersonalAccessTokensPost(context.Background()).CreatePersonalAccessTokenCommand(createPersonalAccessTokenCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensApi.ApiPersonalAccessTokensPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPersonalAccessTokensPost`: PersonalAccessTokenValue
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokensApi.ApiPersonalAccessTokensPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPersonalAccessTokensPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPersonalAccessTokenCommand** | [**CreatePersonalAccessTokenCommand**](CreatePersonalAccessTokenCommand.md) |  | 
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**PersonalAccessTokenValue**](PersonalAccessTokenValue.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

