package services

import (
	"context"
	"fmt"
	"log"

	"tripig/src/config"
	"tripig/src/database"
	"tripig/src/models"

	"googlemaps.github.io/maps"
)

type PlaceTypes []models.PlaceType

func GetPlaceTypeList() []models.PlaceType {
	var placeTypes PlaceTypes
	result := database.Handler.Find(&placeTypes)
	fmt.Print(placeTypes)
	
	if err := result.Error; err != nil {
		fmt.Println(err)
	}

	return placeTypes
}

func NearbySearch(keyword string, pageToken string, placeType string) models.PlacesSearchResponse {
	c, err := maps.NewClient(maps.WithAPIKey(config.App.GoogleMapApiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	getRequest := &maps.GeocodingRequest{
		Address: keyword,
	}
	getResult, err := c.Geocode(context.Background(), getRequest)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := &maps.NearbySearchRequest{
		Location:  &getResult[0].Geometry.Location,
		Radius:    1500,
		Language:  "ja",
		PageToken: pageToken,
	}
	parsePlaceType(placeType, r)

	if len(pageToken) > 0 {
		r.PageToken = pageToken
	}

	resp, err := c.NearbySearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	return *models.FromMapsPlaceSearchResponse(resp)
}

func parsePlaceType(placeType string, r *maps.NearbySearchRequest) {
	if placeType != "" {
		t, err := maps.ParsePlaceType(placeType)
		if err != nil {
			log.Fatalf("fatal error: %s", err)
		}
		r.Type = t
	}
}

func TextSearch(keyword string) []maps.PlacesSearchResult {
	c, err := maps.NewClient(maps.WithAPIKey(config.App.GoogleMapApiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.TextSearchRequest{
		Query: keyword,
	}

	resp, err := c.TextSearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return resp.Results
}
