package models

import "googlemaps.github.io/maps"

type Geometry struct {
	Location Location `json:"location"`
	Viewport Viewport `json:"viewport"`
}

func FromMapsGeometry(apiGeometry maps.AddressGeometry) *Geometry {
	return &Geometry{
		Location: *FromMapsLocation(apiGeometry.Location),
		Viewport: *FromMapsViewport(apiGeometry.Viewport),
	}
}
