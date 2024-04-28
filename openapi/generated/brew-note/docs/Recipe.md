# Recipe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**ExtractionEquipment** | **string** |  | 
**CoffeeType** | **string** |  | 
**Steps** | [**[]RecipeStep**](RecipeStep.md) |  | 

## Methods

### NewRecipe

`func NewRecipe(id float32, extractionEquipment string, coffeeType string, steps []RecipeStep, ) *Recipe`

NewRecipe instantiates a new Recipe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipeWithDefaults

`func NewRecipeWithDefaults() *Recipe`

NewRecipeWithDefaults instantiates a new Recipe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Recipe) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Recipe) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Recipe) SetId(v float32)`

SetId sets Id field to given value.


### GetExtractionEquipment

`func (o *Recipe) GetExtractionEquipment() string`

GetExtractionEquipment returns the ExtractionEquipment field if non-nil, zero value otherwise.

### GetExtractionEquipmentOk

`func (o *Recipe) GetExtractionEquipmentOk() (*string, bool)`

GetExtractionEquipmentOk returns a tuple with the ExtractionEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionEquipment

`func (o *Recipe) SetExtractionEquipment(v string)`

SetExtractionEquipment sets ExtractionEquipment field to given value.


### GetCoffeeType

`func (o *Recipe) GetCoffeeType() string`

GetCoffeeType returns the CoffeeType field if non-nil, zero value otherwise.

### GetCoffeeTypeOk

`func (o *Recipe) GetCoffeeTypeOk() (*string, bool)`

GetCoffeeTypeOk returns a tuple with the CoffeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoffeeType

`func (o *Recipe) SetCoffeeType(v string)`

SetCoffeeType sets CoffeeType field to given value.


### GetSteps

`func (o *Recipe) GetSteps() []RecipeStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Recipe) GetStepsOk() (*[]RecipeStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Recipe) SetSteps(v []RecipeStep)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


