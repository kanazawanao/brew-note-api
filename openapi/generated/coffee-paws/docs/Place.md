# Place

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessStatus** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Lat** | Pointer to **int32** |  | [optional] 
**Lng** | Pointer to **int32** |  | [optional] 
**OpenNow** | Pointer to **bool** |  | [optional] 
**PhotoUrls** | Pointer to **[]string** |  | [optional] 
**PlaceId** | Pointer to **string** |  | [optional] 
**PriceLevel** | Pointer to **int32** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**Types** | Pointer to **[]string** |  | [optional] 
**UserRatingsTotal** | Pointer to **int32** |  | [optional] 
**Vicinity** | Pointer to **string** |  | [optional] 

## Methods

### NewPlace

`func NewPlace() *Place`

NewPlace instantiates a new Place object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceWithDefaults

`func NewPlaceWithDefaults() *Place`

NewPlaceWithDefaults instantiates a new Place object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessStatus

`func (o *Place) GetBusinessStatus() string`

GetBusinessStatus returns the BusinessStatus field if non-nil, zero value otherwise.

### GetBusinessStatusOk

`func (o *Place) GetBusinessStatusOk() (*string, bool)`

GetBusinessStatusOk returns a tuple with the BusinessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessStatus

`func (o *Place) SetBusinessStatus(v string)`

SetBusinessStatus sets BusinessStatus field to given value.

### HasBusinessStatus

`func (o *Place) HasBusinessStatus() bool`

HasBusinessStatus returns a boolean if a field has been set.

### GetIcon

`func (o *Place) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Place) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Place) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Place) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetName

`func (o *Place) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Place) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Place) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Place) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLat

`func (o *Place) GetLat() int32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *Place) GetLatOk() (*int32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *Place) SetLat(v int32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *Place) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *Place) GetLng() int32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *Place) GetLngOk() (*int32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *Place) SetLng(v int32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *Place) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetOpenNow

`func (o *Place) GetOpenNow() bool`

GetOpenNow returns the OpenNow field if non-nil, zero value otherwise.

### GetOpenNowOk

`func (o *Place) GetOpenNowOk() (*bool, bool)`

GetOpenNowOk returns a tuple with the OpenNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenNow

`func (o *Place) SetOpenNow(v bool)`

SetOpenNow sets OpenNow field to given value.

### HasOpenNow

`func (o *Place) HasOpenNow() bool`

HasOpenNow returns a boolean if a field has been set.

### GetPhotoUrls

`func (o *Place) GetPhotoUrls() []string`

GetPhotoUrls returns the PhotoUrls field if non-nil, zero value otherwise.

### GetPhotoUrlsOk

`func (o *Place) GetPhotoUrlsOk() (*[]string, bool)`

GetPhotoUrlsOk returns a tuple with the PhotoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrls

`func (o *Place) SetPhotoUrls(v []string)`

SetPhotoUrls sets PhotoUrls field to given value.

### HasPhotoUrls

`func (o *Place) HasPhotoUrls() bool`

HasPhotoUrls returns a boolean if a field has been set.

### GetPlaceId

`func (o *Place) GetPlaceId() string`

GetPlaceId returns the PlaceId field if non-nil, zero value otherwise.

### GetPlaceIdOk

`func (o *Place) GetPlaceIdOk() (*string, bool)`

GetPlaceIdOk returns a tuple with the PlaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceId

`func (o *Place) SetPlaceId(v string)`

SetPlaceId sets PlaceId field to given value.

### HasPlaceId

`func (o *Place) HasPlaceId() bool`

HasPlaceId returns a boolean if a field has been set.

### GetPriceLevel

`func (o *Place) GetPriceLevel() int32`

GetPriceLevel returns the PriceLevel field if non-nil, zero value otherwise.

### GetPriceLevelOk

`func (o *Place) GetPriceLevelOk() (*int32, bool)`

GetPriceLevelOk returns a tuple with the PriceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevel

`func (o *Place) SetPriceLevel(v int32)`

SetPriceLevel sets PriceLevel field to given value.

### HasPriceLevel

`func (o *Place) HasPriceLevel() bool`

HasPriceLevel returns a boolean if a field has been set.

### GetRating

`func (o *Place) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Place) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Place) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Place) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetTypes

`func (o *Place) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *Place) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *Place) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *Place) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUserRatingsTotal

`func (o *Place) GetUserRatingsTotal() int32`

GetUserRatingsTotal returns the UserRatingsTotal field if non-nil, zero value otherwise.

### GetUserRatingsTotalOk

`func (o *Place) GetUserRatingsTotalOk() (*int32, bool)`

GetUserRatingsTotalOk returns a tuple with the UserRatingsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRatingsTotal

`func (o *Place) SetUserRatingsTotal(v int32)`

SetUserRatingsTotal sets UserRatingsTotal field to given value.

### HasUserRatingsTotal

`func (o *Place) HasUserRatingsTotal() bool`

HasUserRatingsTotal returns a boolean if a field has been set.

### GetVicinity

`func (o *Place) GetVicinity() string`

GetVicinity returns the Vicinity field if non-nil, zero value otherwise.

### GetVicinityOk

`func (o *Place) GetVicinityOk() (*string, bool)`

GetVicinityOk returns a tuple with the Vicinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVicinity

`func (o *Place) SetVicinity(v string)`

SetVicinity sets Vicinity field to given value.

### HasVicinity

`func (o *Place) HasVicinity() bool`

HasVicinity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


