package models

import "googlemaps.github.io/maps"

type Viewport struct {
	Northeast Location `json:"northeast"`
	Southwest Location `json:"southwest"`
}

func FromMapsViewport(apiViewport maps.LatLngBounds) *Viewport {
	return &Viewport{
		Northeast: *FromMapsLocation(apiViewport.NorthEast),
		Southwest: *FromMapsLocation(apiViewport.SouthWest),
	}
}
