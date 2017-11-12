package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	fmt.Println("good to Go")
	mux := http.NewServeMux()
	mux.HandleFunc("/get-weather", WeatherHandler)

}

const (
	apiKey = "42dc862aa7bc58b4d96964d"
	endpoint = "http://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s"
)



type Weather struct {
	description string
	humidity int
	pressure int
	temp int
}
type LatLon struct {
	Latitude string `json:"lat"`
	Longitude string `json:"lon"`
}
func (l *LatLon) Build() string {
	return fmt.Sprintf(endpoint, l.Latitude, l.Longitude, apiKey)
}
func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal(&LatLon{})
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}