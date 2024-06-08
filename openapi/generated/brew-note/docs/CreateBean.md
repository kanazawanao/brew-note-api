# CreateBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductionRegion** | **string** |  | 
**Kind** | **string** |  | 
**RoastLevelId** | **int32** |  | 
**Price** | Pointer to **int32** |  | [optional] 
**Gram** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateBean

`func NewCreateBean(productionRegion string, kind string, roastLevelId int32, ) *CreateBean`

NewCreateBean instantiates a new CreateBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBeanWithDefaults

`func NewCreateBeanWithDefaults() *CreateBean`

NewCreateBeanWithDefaults instantiates a new CreateBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductionRegion

`func (o *CreateBean) GetProductionRegion() string`

GetProductionRegion returns the ProductionRegion field if non-nil, zero value otherwise.

### GetProductionRegionOk

`func (o *CreateBean) GetProductionRegionOk() (*string, bool)`

GetProductionRegionOk returns a tuple with the ProductionRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionRegion

`func (o *CreateBean) SetProductionRegion(v string)`

SetProductionRegion sets ProductionRegion field to given value.


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


### GetRoastLevelId

`func (o *CreateBean) GetRoastLevelId() int32`

GetRoastLevelId returns the RoastLevelId field if non-nil, zero value otherwise.

### GetRoastLevelIdOk

`func (o *CreateBean) GetRoastLevelIdOk() (*int32, bool)`

GetRoastLevelIdOk returns a tuple with the RoastLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoastLevelId

`func (o *CreateBean) SetRoastLevelId(v int32)`

SetRoastLevelId sets RoastLevelId field to given value.


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


