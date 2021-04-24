package http

import (
	"fmt"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
	"io/ioutil"
	"net/http"
	"time"
)

//Get returns response body in bytes with status code or error with status code...
func (dao Service) Get(url string) ([]byte, int, error) {
	log.Info(fmt.Sprintf("making api call to %s", url))

	start := time.Now()
	response, err := dao.HTTPClient.Get(url)
	if err != nil {
		log.Error(fmt.Sprintf("an error occurred while making request to %s with error %s with duration %d ms", url, err.Error(), time.Now().Sub(start).Milliseconds()))
		return nil, http.StatusInternalServerError, err
	}

	log.Info(fmt.Sprintf("recieved response from %s with status code %d with time elapsed %d ms", url, response.StatusCode, time.Since(start).Milliseconds()))

	if response == nil {
		log.Error(fmt.Sprintf("received a nil response from %s", url))
		return nil, http.StatusInternalServerError, fmt.Errorf("an error occurred trying to call http do request")

	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(fmt.Sprintf("an error occurred while reading response body for url %s", url))
		return nil, http.StatusInternalServerError, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		log.Error(fmt.Sprintf("bad response from api call to %s with body %s and status code of %d", url, string(bodyBytes), response.StatusCode))
		return nil, response.StatusCode, fmt.Errorf("an error occurred trying to make api call to %s", url)
	}

	return bodyBytes, response.StatusCode, nil
}
