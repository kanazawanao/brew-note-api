package models

import "googlemaps.github.io/maps"

type Place struct {
	BusinessStatus   string   `json:"businessStatus"`
	Icon             string   `json:"icon"`
	Name             string   `json:"name"`
	Lat              float32  `json:"lat"`
	Lng              float32  `json:"lng"`
	OpenNow          bool     `json:"openNow"`
	Photos           []Photo  `json:"photos"`
	PlaceId          string   `json:"placeId"`
	PriceLevel       int      `json:"priceLevel"`
	Rating           float32  `json:"rating"`
	Types            []string `json:"types"`
	UserRatingsTotal int      `json:"userRatingsTotal"`
	Vicinity         string   `json:"vicinity"`
}

func FromMapsPlace(apiPlace maps.PlacesSearchResult) *Place {
	return &Place{
		BusinessStatus:   apiPlace.BusinessStatus,
		Icon:             apiPlace.Icon,
		Name:             apiPlace.Name,
	  Lat:              float32(apiPlace.Geometry.Location.Lat),
		Lng:              float32(apiPlace.Geometry.Location.Lng),
		OpenNow:          false,
		Photos:           FromMapsPhotos(apiPlace.Photos),
		PlaceId:          apiPlace.PlaceID,
		PriceLevel:       apiPlace.PriceLevel,
		Rating:           apiPlace.Rating,
		Types:            apiPlace.Types,
		UserRatingsTotal: apiPlace.UserRatingsTotal,
		Vicinity:         apiPlace.Vicinity,
	}
}

func FromMapsPlaces(apiPlaces []maps.PlacesSearchResult) []Place {
	result := []Place{}
	for _, apiPlace := range apiPlaces {
		result = append(result, *FromMapsPlace(apiPlace))
	}
	return result
}
