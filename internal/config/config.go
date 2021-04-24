package config

import (
	httpService "city-temp/internal/services/http"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"

	"github.com/joho/godotenv"
)

const (
	envPrefix = "THYCOTIC"
)

//WebServerConfig is used to store the environment and port for the webserver as well as the logging config.
type WebServerConfig struct {
	ENV                    string             `required:"true" split_words:"true"`
	WeatherCityURLTemplate string             `required:"true" split_words:"true"`
	WeatherURLTemplate     string             `required:"true" split_words:"true"`
	CityUrls               string             `required:"true" split_words:"true"`
	MaxProc                int                `required:"true" split_words:"true"`
	HTTP                   httpService.Config `required:"true" split_words:"true"`
}

// FromEnv pulls config from the environment
func FromEnv() (cfg *WebServerConfig, err error) {
	fromFileToEnv()

	cfg = &WebServerConfig{}

	err = envconfig.Process(envPrefix, cfg)

	if err != nil {

		return nil, err
	}

	return cfg, nil
}

func fromFileToEnv() {
	cfgFilename := os.Getenv("ENV_FILE")
	if cfgFilename != "" {
		err := godotenv.Load(cfgFilename)
		if err != nil {
			fmt.Println("ENV_FILE not found. Trying MY_POD_NAMESPACE")
		}
		return
	}

	loc := os.Getenv("ENVIRONMENT_NAME")

	cfgFilename = fmt.Sprintf("./etc/config/config.%s.env", loc)

	err := godotenv.Load(cfgFilename)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s %s", "No config files found to load to env. Defaulting to environment with error %s.", err.Error()))
	}

}
