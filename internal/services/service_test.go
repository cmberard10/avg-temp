package services

import (
	"city-temp/internal/config"
	"city-temp/internal/models"
	mockHTTPDao "city-temp/mock/mockHTTPDAO"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCityData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockHTTP := mockHTTPDao.NewMockDAO(mockCtrl)
	service.httpDao = mockHTTP
	service.webConfig = &config.WebServerConfig{}

	type expect struct {
		resp       *models.City
		StatusCode int
		Err        error
	}

	type mockArgs struct {
		BodyBytes  []byte
		StatusCode int
		Err        error
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
				BodyBytes:  []byte(`{"records": [{}]}`),
				StatusCode: 200,
				Err:        nil,
			},
			Expect: expect{
				resp: &models.City{Records: []models.Record{{}}},
				Err:  nil,
			},
		},
		{
			Name: "api call failed",
			MockArgs: mockArgs{
				StatusCode: 200,
				Err:        fmt.Errorf("err"),
			},
			Expect: expect{
				resp: nil,
				Err:  fmt.Errorf("err"),
			},
		},
		{
			Name: "bad json format",
			MockArgs: mockArgs{
				BodyBytes:  []byte(`{"records": "trst"}`),
				StatusCode: 200,
				Err:        nil,
			},
			Expect: expect{
				resp: nil,
				Err:  fmt.Errorf(""),
			},
		},
	}

	for _, tt := range tests {
		mockHTTP.EXPECT().Get(gomock.Any()).Return(tt.MockArgs.BodyBytes, tt.MockArgs.StatusCode, tt.MockArgs.Err)
		resp, err := service.GetCityData()
		assert.Equal(t, resp, tt.Expect.resp)
		if tt.Expect.Err != nil {
			assert.NotNil(t, err)
		}

	}
}
