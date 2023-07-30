package models

import "googlemaps.github.io/maps"

type PlacesSearchResponse struct {
	Results       []Place `json:"results"`
	NextPageToken string  `json:"nextPageToken"`
}

func FromMapsPlaceSearchResponse(apiResult maps.PlacesSearchResponse) *PlacesSearchResponse {
	return &PlacesSearchResponse{
		Results:       FromMapsPlaces(apiResult.Results),
		NextPageToken: apiResult.NextPageToken,
	}
}
