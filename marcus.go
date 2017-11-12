package main

import (
	"fmt"
)

func main() {
	fmt.Println("good to Go")
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetHandler)
	mux.HandleFunc("/post", PostHandler)

}

const (
	apiKey = ""
	endpoint = "http://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s"
)

type Url struct {
	Endpoint string 
	Lat string
	Lon string
}

func (u *Url) Latitude(latitude string) {
	u.Lat = latitude 
}

func (u *Url) Longitude(longitude string) {
	u.Lon = longitude
}

func (u *Url) Build() string {
	return fmt.Sprintf(endpoint, u.Lat, u.Lon)
}