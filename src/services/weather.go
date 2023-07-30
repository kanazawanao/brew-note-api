package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetCurrentWeather() []byte {
	method := "GET"
	url := "https://api.openweathermap.org/data/2.5/weather"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	apiKey := os.Getenv("OPEN_WEATHER_MAP_API_KEY")
	q := req.URL.Query()
	q.Add("q", "tokyo")
	q.Add("appid", apiKey)
	q.Add("lang", "ja")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return body
}
