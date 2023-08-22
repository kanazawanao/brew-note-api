# PlacesSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | **NullableString** |  | 
**Results** | [**[]Place**](Place.md) |  | 

## Methods

### NewPlacesSearchResponse

`func NewPlacesSearchResponse(nextPageToken NullableString, results []Place, ) *PlacesSearchResponse`

NewPlacesSearchResponse instantiates a new PlacesSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacesSearchResponseWithDefaults

`func NewPlacesSearchResponseWithDefaults() *PlacesSearchResponse`

NewPlacesSearchResponseWithDefaults instantiates a new PlacesSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *PlacesSearchResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *PlacesSearchResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *PlacesSearchResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.


### SetNextPageTokenNil

`func (o *PlacesSearchResponse) SetNextPageTokenNil(b bool)`

 SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil

### UnsetNextPageToken
`func (o *PlacesSearchResponse) UnsetNextPageToken()`

UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil
### GetResults

`func (o *PlacesSearchResponse) GetResults() []Place`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PlacesSearchResponse) GetResultsOk() (*[]Place, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PlacesSearchResponse) SetResults(v []Place)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


