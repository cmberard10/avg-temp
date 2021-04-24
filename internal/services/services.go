package services

import (
	"city-temp/internal/config"
	"city-temp/internal/models"
	httpService "city-temp/internal/services/http"
	"encoding/json"
	"fmt"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
	"time"
)

var service ThycoticService

//ThycoticService type of dao thycotic service utilizes
type ThycoticService struct {
	httpDao   httpService.DAO
	webConfig *config.WebServerConfig
}

//Initialize DAO
func Initialize(webConfig *config.WebServerConfig) *ThycoticService {
	service.webConfig = webConfig
	httpDAO := httpService.Initialize(webConfig.HTTP)
	service.httpDao = httpDAO
	return &service
}

//GetCityData calls api to get 100 cities sorted by population in the us
func (dao ThycoticService) GetCityData() (*models.City, error) {
	bodyBytes, _, err := service.httpDao.Get(service.webConfig.CityUrls)
	if err != nil {
		return nil, err
	}

	var resp models.City
	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		log.Error(fmt.Sprintf("an error occurred whil unmarshaling city object with error %s", err.Error()))
		return nil, err
	}

	return &resp, nil
}

//GetCurrentTemperatureForCoordinates gets weather api id closest to coords passed in and returns the current weather in celsius
func (dao ThycoticService) GetCurrentTemperatureForCoordinates(coord models.Coordinate) (float64, error) {
	weatherCityBodyBytes, _, err := service.httpDao.Get(fmt.Sprintf(service.webConfig.WeatherCityURLTemplate, coord.Longitude, coord.Latitude))
	if err != nil {
		return 0, err
	}

	weatherCity := make([]models.WeatherCity, 0)
	err = json.Unmarshal(weatherCityBodyBytes, &weatherCity)
	if err != nil {
		log.Error(fmt.Sprintf("an error occurred while unmarshaling city object with error %s", err.Error()))
		return 0, err
	}

	if len(weatherCity) == 0 {
		log.Error("did not receive any city info")
		return 0, fmt.Errorf("did not receive any city info")
	}

	//we are using first index because its assumed that woeids are returned in sorted order of closest to lat to farthest
	weatherURLTemplateBodyBytes, _, err := service.httpDao.Get(fmt.Sprintf(service.webConfig.WeatherURLTemplate, weatherCity[0].Woeid, time.Now().Year(), int(time.Now().Month()), time.Now().Day()))
	if err != nil {
		return 0, err
	}

	weather := make([]models.Weather, 0)
	err = json.Unmarshal(weatherURLTemplateBodyBytes, &weather)
	if err != nil {
		log.Error(fmt.Sprintf("an error occurred while unmarshaling city object with error %s", err.Error()))
		return 0, err
	}

	if len(weather) == 0 {
		log.Error("did not receive any city info")
		return 0, fmt.Errorf("did not receive any city info")
	}

	return weather[0].TheTemp, nil
}
