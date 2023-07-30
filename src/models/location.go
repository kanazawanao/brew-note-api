package models

import "googlemaps.github.io/maps"

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func FromMapsLocation(apiLocation maps.LatLng) *Location {
	return &Location{
		apiLocation.Lat,
		apiLocation.Lng,
	}
}
