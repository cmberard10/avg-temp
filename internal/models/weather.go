package models

import "time"

//Coordinate expected object for GetCurrentTemperatureForCoordinates
type Coordinate struct {
	Latitude  string
	Longitude string
}

//WeatherCity response from https://www.metaweather.com/api//location/search?lattlong=%s,%s
type WeatherCity struct {
	Distance     int    `json:"distance"`
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
}

//Weather expected response from https://www.metaweather.com/api/location/%d/%d/%d/%d
type Weather struct {
	ID                   int64     `json:"id"`
	WeatherStateName     string    `json:"weather_state_name"`
	WeatherStateAbbr     string    `json:"weather_state_abbr"`
	WindDirectionCompass string    `json:"wind_direction_compass"`
	Created              time.Time `json:"created"`
	ApplicableDate       string    `json:"applicable_date"`
	MinTemp              float64   `json:"min_temp"`
	MaxTemp              float64   `json:"max_temp"`
	TheTemp              float64   `json:"the_temp"`
	WindSpeed            float64   `json:"wind_speed"`
	WindDirection        float64   `json:"wind_direction"`
	AirPressure          float64   `json:"air_pressure"`
	Humidity             int       `json:"humidity"`
	Visibility           float64   `json:"visibility"`
	Predictability       int       `json:"predictability"`
}
