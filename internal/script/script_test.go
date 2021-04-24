package script

import (
	"city-temp/internal/config"
	"city-temp/internal/models"
	mockSCRIPTDao "city-temp/mock/mockSCRIPTDAO"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/bloom42/libs/rz-go/v2"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
	"testing"
)

func TestGetCityData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockScript := mockSCRIPTDao.NewMockDAO(mockCtrl)
	scriptDAO = mockScript
	copyConfig = &config.WebServerConfig{MaxProc: 1}
	log.SetLogger(log.With(rz.Fields(rz.String("environment", "running-tests"))))

	type expect struct {
		average       float64
		countOfCities int
		err           error
	}

	type GetCityData struct {
		resp *models.City
		err  error
	}

	type GetCurrentTemperatureForCoordinates struct {
		cityTemp float64
		err      error
	}

	type mockArgs struct {
		GetCityData                         GetCityData
		GetCurrentTemperatureForCoordinates GetCurrentTemperatureForCoordinates
	}

	type test struct {
		Name     string
		MockArgs mockArgs
		Expect   expect
	}

	tests := []test{
		{
			Name: "success",
			MockArgs: mockArgs{
				GetCityData: GetCityData{
					resp: &models.City{Records: []models.Record{
						{},
					}},
					err: nil,
				},
				GetCurrentTemperatureForCoordinates: GetCurrentTemperatureForCoordinates{
					cityTemp: 5,
					err:      nil,
				},
			},
			Expect: expect{
				average:       5,
				countOfCities: 1,
				err:           nil,
			},
		},
		{
			Name: "get city data error",
			MockArgs: mockArgs{
				GetCityData: GetCityData{
					err: fmt.Errorf("err"),
				},
				GetCurrentTemperatureForCoordinates: GetCurrentTemperatureForCoordinates{
					cityTemp: 5,
					err:      nil,
				},
			},
			Expect: expect{
				err: fmt.Errorf("err"),
			},
		},
		{
			Name: "no cities found",
			MockArgs: mockArgs{
				GetCityData: GetCityData{
					resp: &models.City{},
				},
				GetCurrentTemperatureForCoordinates: GetCurrentTemperatureForCoordinates{
					cityTemp: 5,
					err:      nil,
				},
			},
			Expect: expect{
				err: fmt.Errorf("no city records found"),
			},
		},
		{
			Name: "error on get temp",
			MockArgs: mockArgs{
				GetCityData: GetCityData{
					resp: &models.City{Records: []models.Record{
						{},
					}},
					err: nil,
				},
				GetCurrentTemperatureForCoordinates: GetCurrentTemperatureForCoordinates{
					cityTemp: 0,
					err:      fmt.Errorf("err"),
				},
			},
			Expect: expect{
				average:       0,
				countOfCities: 0,
				err:           nil,
			},
		},
	}

	for _, tt := range tests {
		log.Info(tt.Name)
		mockScript.EXPECT().GetCityData().Return(tt.MockArgs.GetCityData.resp, tt.MockArgs.GetCityData.err)
		if tt.MockArgs.GetCityData.err == nil && tt.MockArgs.GetCityData.resp != nil && len(tt.MockArgs.GetCityData.resp.Records) > 0 {
			mockScript.EXPECT().GetCurrentTemperatureForCoordinates(gomock.Any()).Return(tt.MockArgs.GetCurrentTemperatureForCoordinates.cityTemp, tt.MockArgs.GetCurrentTemperatureForCoordinates.err)
		}

		avg, count, err := RunScript()
		assert.Equal(t, avg, tt.Expect.average)
		assert.Equal(t, count, tt.Expect.countOfCities)
		if tt.Expect.err != nil {
			assert.NotNil(t, err)
		}

	}
}
