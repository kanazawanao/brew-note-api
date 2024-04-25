# \BrewNtoeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProfile**](BrewNtoeApi.md#GetProfile) | **Get** /me | Get Profile
[**GetUsers**](BrewNtoeApi.md#GetUsers) | **Get** /users | Get Users



## GetProfile

> User GetProfile(ctx).Execute()

Get Profile



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
    resp, r, err := apiClient.BrewNtoeApi.GetProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrewNtoeApi.GetProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfile`: User
    fmt.Fprintf(os.Stdout, "Response from `BrewNtoeApi.GetProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> []User GetUsers(ctx).Execute()

Get Users



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
    resp, r, err := apiClient.BrewNtoeApi.GetUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrewNtoeApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `BrewNtoeApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

