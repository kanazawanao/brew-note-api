package models

import (
	"tripig/src/config"

	"googlemaps.github.io/maps"
)


func FromMapsPhoto(apiPhoto maps.Photo) string {
	return "https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference=" + apiPhoto.PhotoReference + "&key="+ config.App.GoogleMapApiKey
}

func FromMapsPhotos(apiPhotos []maps.Photo) []string {
	result := []string{}
	for _, apiPhoto := range apiPhotos {
		result = append(result, FromMapsPhoto(apiPhoto))
	}
	return result
}
