package models

import "googlemaps.github.io/maps"

type Photo struct {
	Height           int      `json:"height"`
	Width            int      `json:"width"`
	HtmlAttributions []string `json:"htmlAttributions"`
	PhotoReference   string   `json:"photoReference"`
}

func FromMapsPhoto(apiPhoto maps.Photo) *Photo {
	return &Photo{
		Height:           apiPhoto.Height,
		Width:            apiPhoto.Width,
		HtmlAttributions: apiPhoto.HTMLAttributions,
		PhotoReference:   apiPhoto.PhotoReference,
	}
}

func FromMapsPhotos(apiPhotos []maps.Photo) []Photo {
	result := []Photo{}
	for _, apiPhoto := range apiPhotos {
		result = append(result, *FromMapsPhoto(apiPhoto))
	}
	return result
}
