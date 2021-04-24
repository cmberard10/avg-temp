package services

import "city-temp/internal/models"

//DAO services dao
type DAO interface {
	GetCityData() (*models.City, error)
	GetCurrentTemperatureForCoordinates(coord models.Coordinate) (float64, error)
}
