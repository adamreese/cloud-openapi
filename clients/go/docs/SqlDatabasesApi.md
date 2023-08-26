# \SqlDatabasesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSqlDatabasesCreatePost**](SqlDatabasesApi.md#ApiSqlDatabasesCreatePost) | **Post** /api/sql-databases/create | 
[**ApiSqlDatabasesDelete**](SqlDatabasesApi.md#ApiSqlDatabasesDelete) | **Delete** /api/sql-databases | 
[**ApiSqlDatabasesExecutePost**](SqlDatabasesApi.md#ApiSqlDatabasesExecutePost) | **Post** /api/sql-databases/execute | 
[**ApiSqlDatabasesGet**](SqlDatabasesApi.md#ApiSqlDatabasesGet) | **Get** /api/sql-databases | 



## ApiSqlDatabasesCreatePost

> ApiSqlDatabasesCreatePost(ctx).CreateSqlDatabaseCommand(createSqlDatabaseCommand).ApiVersion(apiVersion).Execute()



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
    createSqlDatabaseCommand := *openapiclient.NewCreateSqlDatabaseCommand("Name_example") // CreateSqlDatabaseCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SqlDatabasesApi.ApiSqlDatabasesCreatePost(context.Background()).CreateSqlDatabaseCommand(createSqlDatabaseCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlDatabasesApi.ApiSqlDatabasesCreatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSqlDatabasesCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSqlDatabaseCommand** | [**CreateSqlDatabaseCommand**](CreateSqlDatabaseCommand.md) |  | 
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


## ApiSqlDatabasesDelete

> ApiSqlDatabasesDelete(ctx).DeleteSqlDatabaseCommand(deleteSqlDatabaseCommand).ApiVersion(apiVersion).Execute()



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
    deleteSqlDatabaseCommand := *openapiclient.NewDeleteSqlDatabaseCommand("Name_example") // DeleteSqlDatabaseCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SqlDatabasesApi.ApiSqlDatabasesDelete(context.Background()).DeleteSqlDatabaseCommand(deleteSqlDatabaseCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlDatabasesApi.ApiSqlDatabasesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSqlDatabasesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSqlDatabaseCommand** | [**DeleteSqlDatabaseCommand**](DeleteSqlDatabaseCommand.md) |  | 
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


## ApiSqlDatabasesExecutePost

> ApiSqlDatabasesExecutePost(ctx).ExecuteSqlStatementCommand(executeSqlStatementCommand).ApiVersion(apiVersion).Execute()



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
    executeSqlStatementCommand := *openapiclient.NewExecuteSqlStatementCommand("Database_example", false, "Statement_example") // ExecuteSqlStatementCommand | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SqlDatabasesApi.ApiSqlDatabasesExecutePost(context.Background()).ExecuteSqlStatementCommand(executeSqlStatementCommand).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlDatabasesApi.ApiSqlDatabasesExecutePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSqlDatabasesExecutePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executeSqlStatementCommand** | [**ExecuteSqlStatementCommand**](ExecuteSqlStatementCommand.md) |  | 
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


## ApiSqlDatabasesGet

> DatabasesList ApiSqlDatabasesGet(ctx).GetSqlDatabasesQuery(getSqlDatabasesQuery).ApiVersion(apiVersion).Execute()



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
    getSqlDatabasesQuery := *openapiclient.NewGetSqlDatabasesQuery() // GetSqlDatabasesQuery | 
    apiVersion := "apiVersion_example" // string | The requested API version (optional) (default to "1.0")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SqlDatabasesApi.ApiSqlDatabasesGet(context.Background()).GetSqlDatabasesQuery(getSqlDatabasesQuery).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SqlDatabasesApi.ApiSqlDatabasesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSqlDatabasesGet`: DatabasesList
    fmt.Fprintf(os.Stdout, "Response from `SqlDatabasesApi.ApiSqlDatabasesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSqlDatabasesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSqlDatabasesQuery** | [**GetSqlDatabasesQuery**](GetSqlDatabasesQuery.md) |  | 
 **apiVersion** | **string** | The requested API version | [default to &quot;1.0&quot;]

### Return type

[**DatabasesList**](DatabasesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

