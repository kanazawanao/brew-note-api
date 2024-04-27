# Bean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**BeanId** | **string** |  | 
**ProductionArea** | **string** |  | 
**Kind** | **string** |  | 
**RoastLevel** | Pointer to [**RoastLevel**](RoastLevel.md) |  | [optional] 
**Price** | **NullableFloat64** |  | 

## Methods

### NewBean

`func NewBean(id string, beanId string, productionArea string, kind string, price NullableFloat64, ) *Bean`

NewBean instantiates a new Bean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeanWithDefaults

`func NewBeanWithDefaults() *Bean`

NewBeanWithDefaults instantiates a new Bean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bean) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bean) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bean) SetId(v string)`

SetId sets Id field to given value.


### GetBeanId

`func (o *Bean) GetBeanId() string`

GetBeanId returns the BeanId field if non-nil, zero value otherwise.

### GetBeanIdOk

`func (o *Bean) GetBeanIdOk() (*string, bool)`

GetBeanIdOk returns a tuple with the BeanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeanId

`func (o *Bean) SetBeanId(v string)`

SetBeanId sets BeanId field to given value.


### GetProductionArea

`func (o *Bean) GetProductionArea() string`

GetProductionArea returns the ProductionArea field if non-nil, zero value otherwise.

### GetProductionAreaOk

`func (o *Bean) GetProductionAreaOk() (*string, bool)`

GetProductionAreaOk returns a tuple with the ProductionArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionArea

`func (o *Bean) SetProductionArea(v string)`

SetProductionArea sets ProductionArea field to given value.


### GetKind

`func (o *Bean) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Bean) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Bean) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRoastLevel

`func (o *Bean) GetRoastLevel() RoastLevel`

GetRoastLevel returns the RoastLevel field if non-nil, zero value otherwise.

### GetRoastLevelOk

`func (o *Bean) GetRoastLevelOk() (*RoastLevel, bool)`

GetRoastLevelOk returns a tuple with the RoastLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoastLevel

`func (o *Bean) SetRoastLevel(v RoastLevel)`

SetRoastLevel sets RoastLevel field to given value.

### HasRoastLevel

`func (o *Bean) HasRoastLevel() bool`

HasRoastLevel returns a boolean if a field has been set.

### GetPrice

`func (o *Bean) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Bean) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Bean) SetPrice(v float64)`

SetPrice sets Price field to given value.


### SetPriceNil

`func (o *Bean) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *Bean) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


