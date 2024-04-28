# \BrewNoteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBean**](BrewNoteApi.md#CreateBean) | **Post** /beans | Create Bean
[**CreateRecipe**](BrewNoteApi.md#CreateRecipe) | **Post** /recipes | Post Recipe
[**GetBeans**](BrewNoteApi.md#GetBeans) | **Get** /beans | Get Beans
[**GetRecipes**](BrewNoteApi.md#GetRecipes) | **Get** /recipes | Get Recipes
[**GetRoastLevels**](BrewNoteApi.md#GetRoastLevels) | **Get** /roast-levels | Get Roast Levels



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


## CreateRecipe

> Recipe CreateRecipe(ctx).CreateRecipe(createRecipe).Execute()

Post Recipe



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
    createRecipe := *openapiclient.NewCreateRecipe("ExtractionEquipment_example", "CoffeeType_example", []openapiclient.RecipeStep{*openapiclient.NewRecipeStep(float32(123), float32(123), "Memo_example", float32(123), float32(123))}) // CreateRecipe | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrewNoteApi.CreateRecipe(context.Background()).CreateRecipe(createRecipe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrewNoteApi.CreateRecipe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRecipe`: Recipe
    fmt.Fprintf(os.Stdout, "Response from `BrewNoteApi.CreateRecipe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecipeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRecipe** | [**CreateRecipe**](CreateRecipe.md) |  | 

### Return type

[**Recipe**](Recipe.md)

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


## GetRecipes

> []Recipe GetRecipes(ctx).Execute()

Get Recipes



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
    resp, r, err := apiClient.BrewNoteApi.GetRecipes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrewNoteApi.GetRecipes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecipes`: []Recipe
    fmt.Fprintf(os.Stdout, "Response from `BrewNoteApi.GetRecipes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecipesRequest struct via the builder pattern


### Return type

[**[]Recipe**](Recipe.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoastLevels

> []RoastLevel GetRoastLevels(ctx).Execute()

Get Roast Levels



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
    resp, r, err := apiClient.BrewNoteApi.GetRoastLevels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrewNoteApi.GetRoastLevels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoastLevels`: []RoastLevel
    fmt.Fprintf(os.Stdout, "Response from `BrewNoteApi.GetRoastLevels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoastLevelsRequest struct via the builder pattern


### Return type

[**[]RoastLevel**](RoastLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

