package script

import (
	"city-temp/internal/config"
	"city-temp/internal/models"
	"city-temp/internal/services"
	"fmt"
	"github.com/remeh/sizedwaitgroup"
	"gitlab.com/bloom42/libs/rz-go/v2"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
)

var scriptDAO services.DAO
var copyConfig *config.WebServerConfig

//Initialize services dao
func Initialize() error {
	//get env vars
	webServerConfig, err := config.FromEnv()
	if err != nil {
		return err
	}
	copyConfig = webServerConfig

	log.SetLogger(log.With(rz.Fields(
		rz.String("environment", webServerConfig.ENV),
	)))

	scriptDAO = services.Initialize(webServerConfig)
	return nil
}

//RunScript to get city data and get average of temp in Celsius
func RunScript() (float64, int, error) {
	city, err := scriptDAO.GetCityData()
	if err != nil {
		return 0, 0, err
	}

	if len(city.Records) == 0 {
		log.Error("no city records found")
		return 0, 0, fmt.Errorf("no city records found")
	}

	totalTemp := 0.00
	totalCitiesCalculated := 0

	wg := sizedwaitgroup.New(copyConfig.MaxProc)
	for i := 0; i < len(city.Records); i++ {
		copyI := i
		wg.Add()
		go func(i int) {
			defer wg.Done()
			log.Info(fmt.Sprintf("searching city at index %d", i))
			cityTemp, err := scriptDAO.GetCurrentTemperatureForCoordinates(models.Coordinate{
				Latitude:  city.Records[i].Fields.Longitude,
				Longitude: city.Records[i].Fields.Latitude,
			})
			if err == nil {
				totalTemp += cityTemp
				totalCitiesCalculated++
			}
		}(copyI)
	}

	wg.Wait()

	//assuming this is celsius
	fmt.Printf("average temp %f for %d cities", totalTemp/float64(totalCitiesCalculated), totalCitiesCalculated)

	if totalCitiesCalculated == 0 {
		return 0, 0, nil
	}

	return totalTemp / float64(totalCitiesCalculated), totalCitiesCalculated, nil
}
