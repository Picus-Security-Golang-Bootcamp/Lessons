package models

type Location struct {
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
	CityName    string `json:"cityName"`
	IATACode    string `json:"iataCode"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type LocationSlice []Location
