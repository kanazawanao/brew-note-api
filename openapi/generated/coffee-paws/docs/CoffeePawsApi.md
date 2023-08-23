# \CoffeePawsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBean**](CoffeePawsApi.md#CreateBean) | **Post** /coffee-paws/stores/{id}/beans | Create Bean
[**CreateStore**](CoffeePawsApi.md#CreateStore) | **Post** /coffee-paws/stores | Create Store
[**GetBeans**](CoffeePawsApi.md#GetBeans) | **Get** /coffee-paws/stores/{id}/beans | Get Beans
[**GetPlaceTypes**](CoffeePawsApi.md#GetPlaceTypes) | **Get** /coffee-paws/places/types | List Place Type
[**GetStore**](CoffeePawsApi.md#GetStore) | **Get** /coffee-paws/stores/{id} | Get a store by ID
[**GetStores**](CoffeePawsApi.md#GetStores) | **Get** /coffee-paws/stores | Get Stores
[**GetUser**](CoffeePawsApi.md#GetUser) | **Get** /coffee-paws/users/{id} | Get a user by ID
[**GetUsers**](CoffeePawsApi.md#GetUsers) | **Get** /coffee-paws/users | Get Users
[**SearchNearby**](CoffeePawsApi.md#SearchNearby) | **Get** /coffee-paws/places/nearby | /places/nearby



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
    openapiclient "github.com/coffee-paws/api"
)

func main() {
    createBean := *openapiclient.NewCreateBean("ProductionArea_example", "PlantationName_example", "Kind_example", "RoastLevel_example", "Price_example") // CreateBean | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoffeePawsApi.CreateBean(context.Background()).CreateBean(createBean).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoffeePawsApi.CreateBean``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBean`: Bean
    fmt.Fprintf(os.Stdout, "Response from `CoffeePawsApi.CreateBean`: %v\n", resp)
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


## CreateStore

> Store CreateStore(ctx).CreateStore(createStore).Execute()

Create Store



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/coffee-paws/api"
)

func main() {
    createStore := *openapiclient.NewCreateStore("Name_example", "StoreType_example", "Address_example", "Url_example") // CreateStore | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoffeePawsApi.CreateStore(context.Background()).CreateStore(createStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoffeePawsApi.CreateStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStore`: Store
    fmt.Fprintf(os.Stdout, "Response from `CoffeePawsApi.CreateStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStore** | [**CreateStore**](CreateStore.md) |  | 

### Return type

[**Store**](Store.md)

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
    openapiclient "github.com/coffee-paws/api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoffeePawsApi.GetBeans(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoffeePawsApi.GetBeans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBeans`: []Bean
    fmt.Fprintf(os.Stdout, "Response from `CoffeePawsApi.GetBeans`: %v\n", resp)
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


## GetPlaceTypes

> []PlaceType GetPlaceTypes(ctx).Execute()

List Place Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/coffee-paws/api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoffeePawsApi.GetPlaceTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoffeePawsApi.GetPlaceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlaceTypes`: []PlaceType
    fmt.Fprintf(os.Stdout, "Response from `CoffeePawsApi.GetPlaceTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaceTypesRequest struct via the builder pattern


### Return type

[**[]PlaceType**](PlaceType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStore

> Store GetStore(ctx, id).Execute()

Get a store by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/coffee-paws/api"
)

func main() {
    id := "id_example" // string | store id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoffeePawsApi.GetStore(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoffeePawsApi.GetStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStore`: Store
    fmt.Fprintf(os.Stdout, "Response from `CoffeePawsApi.GetStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | store id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Store**](Store.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStores

> []Store GetStores(ctx).Execute()

Get Stores



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/coffee-paws/api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoffeePawsApi.GetStores(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoffeePawsApi.GetStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStores`: []Store
    fmt.Fprintf(os.Stdout, "Response from `CoffeePawsApi.GetStores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoresRequest struct via the builder pattern


### Return type

[**[]Store**](Store.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, id).Execute()

Get a user by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/coffee-paws/api"
)

func main() {
    id := "id_example" // string | user id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoffeePawsApi.GetUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoffeePawsApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: User
    fmt.Fprintf(os.Stdout, "Response from `CoffeePawsApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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
    openapiclient "github.com/coffee-paws/api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoffeePawsApi.GetUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoffeePawsApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `CoffeePawsApi.GetUsers`: %v\n", resp)
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


## SearchNearby

> PlacesSearchResponse SearchNearby(ctx).Keyword(keyword).PlaceType(placeType).PageToken(pageToken).Execute()

/places/nearby



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/coffee-paws/api"
)

func main() {
    keyword := "keyword_example" // string | keyword
    placeType := "placeType_example" // string | placeType (optional)
    pageToken := "pageToken_example" // string | pageToken (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoffeePawsApi.SearchNearby(context.Background()).Keyword(keyword).PlaceType(placeType).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoffeePawsApi.SearchNearby``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchNearby`: PlacesSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `CoffeePawsApi.SearchNearby`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchNearbyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyword** | **string** | keyword | 
 **placeType** | **string** | placeType | 
 **pageToken** | **string** | pageToken | 

### Return type

[**PlacesSearchResponse**](PlacesSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

