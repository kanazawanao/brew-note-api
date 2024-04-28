# CreateRecipe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtractionEquipment** | **string** |  | 
**CoffeeType** | **string** |  | 
**Steps** | [**[]RecipeStep**](RecipeStep.md) |  | 

## Methods

### NewCreateRecipe

`func NewCreateRecipe(extractionEquipment string, coffeeType string, steps []RecipeStep, ) *CreateRecipe`

NewCreateRecipe instantiates a new CreateRecipe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecipeWithDefaults

`func NewCreateRecipeWithDefaults() *CreateRecipe`

NewCreateRecipeWithDefaults instantiates a new CreateRecipe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtractionEquipment

`func (o *CreateRecipe) GetExtractionEquipment() string`

GetExtractionEquipment returns the ExtractionEquipment field if non-nil, zero value otherwise.

### GetExtractionEquipmentOk

`func (o *CreateRecipe) GetExtractionEquipmentOk() (*string, bool)`

GetExtractionEquipmentOk returns a tuple with the ExtractionEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractionEquipment

`func (o *CreateRecipe) SetExtractionEquipment(v string)`

SetExtractionEquipment sets ExtractionEquipment field to given value.


### GetCoffeeType

`func (o *CreateRecipe) GetCoffeeType() string`

GetCoffeeType returns the CoffeeType field if non-nil, zero value otherwise.

### GetCoffeeTypeOk

`func (o *CreateRecipe) GetCoffeeTypeOk() (*string, bool)`

GetCoffeeTypeOk returns a tuple with the CoffeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoffeeType

`func (o *CreateRecipe) SetCoffeeType(v string)`

SetCoffeeType sets CoffeeType field to given value.


### GetSteps

`func (o *CreateRecipe) GetSteps() []RecipeStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CreateRecipe) GetStepsOk() (*[]RecipeStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CreateRecipe) SetSteps(v []RecipeStep)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


