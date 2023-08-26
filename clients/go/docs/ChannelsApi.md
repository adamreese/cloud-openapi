# \ChannelsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiChannelsChannelIdDesiredStatusPut**](ChannelsApi.md#ApiChannelsChannelIdDesiredStatusPut) | **Put** /api/channels/{channelId}/desired-status | 
[**ApiChannelsGet**](ChannelsApi.md#ApiChannelsGet) | **Get** /api/channels | 
[**ApiChannelsIdDelete**](ChannelsApi.md#ApiChannelsIdDelete) | **Delete** /api/channels/{id} | 
[**ApiChannelsIdGet**](ChannelsApi.md#ApiChannelsIdGet) | **Get** /api/channels/{id} | 
[**ApiChannelsIdHealthzGet**](ChannelsApi.md#ApiChannelsIdHealthzGet) | **Get** /api/channels/{id}/healthz | 
[**ApiChannelsIdLogsGet**](ChannelsApi.md#ApiChannelsIdLogsGet) | **Get** /api/channels/{id}/logs | 
[**ApiChannelsIdPatch**](ChannelsApi.md#ApiChannelsIdPatch) | **Patch** /api/channels/{id} | 
[**ApiChannelsIdPut**](ChannelsApi.md#ApiChannelsIdPut) | **Put** /api/channels/{id} | 
[**ApiChannelsPost**](ChannelsApi.md#ApiChannelsPost) | **Post** /api/channels | 



## ApiChannelsChannelIdDesiredStatusPut

> ApiChannelsChannelIdDesiredStatusPut(ctx, channelId).UpdateDesiredStatusCommand(updateDesiredStatusCommand).ApiVersion(apiVersion).Execute()



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
    channelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    updateDesiredStatusCommand := *openapiclient.NewUpdateDesiredStatusCommand("ChannelId_example", openapiclient.DesiredStatus("Running")) // UpdateDesiredStatusCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ChannelsApi.ApiChannelsChannelIdDesiredStatusPut(context.Background(), channelId).UpdateDesiredStatusCommand(updateDesiredStatusCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.ApiChannelsChannelIdDesiredStatusPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiChannelsChannelIdDesiredStatusPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDesiredStatusCommand** | [**UpdateDesiredStatusCommand**](UpdateDesiredStatusCommand.md) |  | 
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


## ApiChannelsGet

> ChannelItemPage ApiChannelsGet(ctx).SearchText(searchText).PageIndex(pageIndex).PageSize(pageSize).SortBy(sortBy).IsSortedAscending(isSortedAscending).ApiVersion(apiVersion).Execute()



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
    sortBy := "sortBy_example" // string |  (optional) (default to "Name")
    isSortedAscending := true // bool |  (optional) (default to true)
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelsApi.ApiChannelsGet(context.Background()).SearchText(searchText).PageIndex(pageIndex).PageSize(pageSize).SortBy(sortBy).IsSortedAscending(isSortedAscending).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.ApiChannelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiChannelsGet`: ChannelItemPage
    fmt.Fprintf(os.Stdout, "Response from `ChannelsApi.ApiChannelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiChannelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchText** | **string** |  | [default to &quot;&quot;]
 **pageIndex** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 50]
 **sortBy** | **string** |  | [default to &quot;Name&quot;]
 **isSortedAscending** | **bool** |  | [default to true]
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**ChannelItemPage**](ChannelItemPage.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiChannelsIdDelete

> ApiChannelsIdDelete(ctx, id).ApiVersion(apiVersion).Execute()



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
    r, err := apiClient.ChannelsApi.ApiChannelsIdDelete(context.Background(), id).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.ApiChannelsIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiChannelsIdDeleteRequest struct via the builder pattern


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


## ApiChannelsIdGet

> ChannelItem ApiChannelsIdGet(ctx, id).ApiVersion(apiVersion).Execute()



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
    resp, r, err := apiClient.ChannelsApi.ApiChannelsIdGet(context.Background(), id).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.ApiChannelsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiChannelsIdGet`: ChannelItem
    fmt.Fprintf(os.Stdout, "Response from `ChannelsApi.ApiChannelsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiChannelsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**ChannelItem**](ChannelItem.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiChannelsIdHealthzGet

> HealthCheckResult ApiChannelsIdHealthzGet(ctx, id).ApiVersion(apiVersion).Execute()



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
    resp, r, err := apiClient.ChannelsApi.ApiChannelsIdHealthzGet(context.Background(), id).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.ApiChannelsIdHealthzGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiChannelsIdHealthzGet`: HealthCheckResult
    fmt.Fprintf(os.Stdout, "Response from `ChannelsApi.ApiChannelsIdHealthzGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiChannelsIdHealthzGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**HealthCheckResult**](HealthCheckResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiChannelsIdLogsGet

> GetChannelLogsVm ApiChannelsIdLogsGet(ctx, id).Max(max).ApiVersion(apiVersion).Execute()



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
    max := int32(56) // int32 |  (optional)
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelsApi.ApiChannelsIdLogsGet(context.Background(), id).Max(max).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.ApiChannelsIdLogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiChannelsIdLogsGet`: GetChannelLogsVm
    fmt.Fprintf(os.Stdout, "Response from `ChannelsApi.ApiChannelsIdLogsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiChannelsIdLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **max** | **int32** |  | 
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**GetChannelLogsVm**](GetChannelLogsVm.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiChannelsIdPatch

> ApiChannelsIdPatch(ctx, id).PatchChannelCommand(patchChannelCommand).ApiVersion(apiVersion).Execute()



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
    patchChannelCommand := *openapiclient.NewPatchChannelCommand() // PatchChannelCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ChannelsApi.ApiChannelsIdPatch(context.Background(), id).PatchChannelCommand(patchChannelCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.ApiChannelsIdPatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiChannelsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchChannelCommand** | [**PatchChannelCommand**](PatchChannelCommand.md) |  | 
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


## ApiChannelsIdPut

> ApiChannelsIdPut(ctx, id).UpdateChannelCommand(updateChannelCommand).ApiVersion(apiVersion).Execute()



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
    updateChannelCommand := *openapiclient.NewUpdateChannelCommand("Id_example", "Name_example", openapiclient.ChannelRevisionSelectionStrategy("UseRangeRule")) // UpdateChannelCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ChannelsApi.ApiChannelsIdPut(context.Background(), id).UpdateChannelCommand(updateChannelCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.ApiChannelsIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiChannelsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateChannelCommand** | [**UpdateChannelCommand**](UpdateChannelCommand.md) |  | 
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


## ApiChannelsPost

> string ApiChannelsPost(ctx).CreateChannelCommand(createChannelCommand).ApiVersion(apiVersion).Execute()



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
    createChannelCommand := *openapiclient.NewCreateChannelCommand("AppId_example", "Name_example", openapiclient.ChannelRevisionSelectionStrategy("UseRangeRule")) // CreateChannelCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelsApi.ApiChannelsPost(context.Background()).CreateChannelCommand(createChannelCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelsApi.ApiChannelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiChannelsPost`: string
    fmt.Fprintf(os.Stdout, "Response from `ChannelsApi.ApiChannelsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiChannelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createChannelCommand** | [**CreateChannelCommand**](CreateChannelCommand.md) |  | 
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

**string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

