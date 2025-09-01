package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GoldenHour(latitude float64, longitude float64, date string) string {
	// TODO: need to construct the URL from the latitude and longitude
	// TODO: move this somewhere else
	var url = "https://api.sunrise-sunset.org/json?lat=" + strconv.FormatFloat(latitude, 'f', -1, 64) + "&lng=" + strconv.FormatFloat(longitude, 'f', -1, 64)
	if date != "" {
		url += "&date=" + date
	}
	resp, err := http.Get(url)

	// TODO: probably better to return an error instead of logging
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // we read the body
	// TODO: probably better to return an error instead of logging
	if err != nil {
		log.Fatalln(err)
	}

	var bodyMap map[string]interface{}

	err = json.Unmarshal(body, &bodyMap)
	// TODO: probably better to return an error instead of logging
	if err != nil {
		log.Fatalln(err)
	}
	sunset, ok := bodyMap["results"].(map[string]interface{})["sunset"].(string)
	// TODO: probably better to return an error instead of logging
	if !ok {
		log.Fatalln("sunset not found")
	}

	sunsetTime, err := time.Parse("3:04:05 PM", sunset)
	// TODO: probably better to return an error instead of logging
	if err != nil {
		log.Fatalln(err)
	}
	sunsetTime = sunsetTime.Add(-1 * time.Hour) // we want to go back one hour to calcultate the golden hour

	return sunsetTime.Format("3:04:05 PM")
}