package services

import (
	"context"
	"log"
	"os"

	"tripig/src/config"
	"tripig/src/models"

	"googlemaps.github.io/maps"
)

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

// TODO: Query(検索条件) を引数に検索する
func TextSearch() []maps.PlacesSearchResult {
	apiKey := os.Getenv("GOOGLE_MAP_API_KEY")
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.TextSearchRequest{
		Query: "東京タワー",
	}

	resp, err := c.TextSearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return resp.Results
}
