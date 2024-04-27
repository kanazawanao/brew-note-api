# CreateBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductionArea** | **string** |  | 
**Kind** | **string** |  | 
**RoastId** | **int32** |  | 
**Price** | Pointer to **int32** |  | [optional] 
**Gram** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateBean

`func NewCreateBean(productionArea string, kind string, roastId int32, ) *CreateBean`

NewCreateBean instantiates a new CreateBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBeanWithDefaults

`func NewCreateBeanWithDefaults() *CreateBean`

NewCreateBeanWithDefaults instantiates a new CreateBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductionArea

`func (o *CreateBean) GetProductionArea() string`

GetProductionArea returns the ProductionArea field if non-nil, zero value otherwise.

### GetProductionAreaOk

`func (o *CreateBean) GetProductionAreaOk() (*string, bool)`

GetProductionAreaOk returns a tuple with the ProductionArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionArea

`func (o *CreateBean) SetProductionArea(v string)`

SetProductionArea sets ProductionArea field to given value.


### GetKind

`func (o *CreateBean) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateBean) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateBean) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRoastId

`func (o *CreateBean) GetRoastId() int32`

GetRoastId returns the RoastId field if non-nil, zero value otherwise.

### GetRoastIdOk

`func (o *CreateBean) GetRoastIdOk() (*int32, bool)`

GetRoastIdOk returns a tuple with the RoastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoastId

`func (o *CreateBean) SetRoastId(v int32)`

SetRoastId sets RoastId field to given value.


### GetPrice

`func (o *CreateBean) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateBean) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateBean) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CreateBean) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetGram

`func (o *CreateBean) GetGram() int32`

GetGram returns the Gram field if non-nil, zero value otherwise.

### GetGramOk

`func (o *CreateBean) GetGramOk() (*int32, bool)`

GetGramOk returns a tuple with the Gram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGram

`func (o *CreateBean) SetGram(v int32)`

SetGram sets Gram field to given value.

### HasGram

`func (o *CreateBean) HasGram() bool`

HasGram returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


