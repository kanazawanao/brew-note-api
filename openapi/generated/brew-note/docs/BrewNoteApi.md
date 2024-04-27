# \BrewNoteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBean**](BrewNoteApi.md#CreateBean) | **Post** /beans | Create Bean
[**GetBeans**](BrewNoteApi.md#GetBeans) | **Get** /beans | Get Beans



## CreateBean

> Bean CreateBean(ctx).CreateBean(createBean).Execute()

Create Bean



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/brew-note/api"
)

func main() {
    createBean := *openapiclient.NewCreateBean("ProductionArea_example", "Kind_example", int32(123)) // CreateBean | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrewNoteApi.CreateBean(context.Background()).CreateBean(createBean).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrewNoteApi.CreateBean``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBean`: Bean
    fmt.Fprintf(os.Stdout, "Response from `BrewNoteApi.CreateBean`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBeanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBean** | [**CreateBean**](CreateBean.md) |  | 

### Return type

[**Bean**](Bean.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBeans

> []Bean GetBeans(ctx).Execute()

Get Beans



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/brew-note/api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrewNoteApi.GetBeans(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrewNoteApi.GetBeans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBeans`: []Bean
    fmt.Fprintf(os.Stdout, "Response from `BrewNoteApi.GetBeans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBeansRequest struct via the builder pattern


### Return type

[**[]Bean**](Bean.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

