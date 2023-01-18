package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"unitesttraining/src/api/domain/locations"
	"unitesttraining/src/api/utils/errors"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountriesNotFound(t *testing.T) {
	//Init
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"status": 404, "error": "not_found", "message": "no country with id AR"}`,
	})

	//Execution
	response, err := http.Get("http://localhost:3000/locations/countries/AR")

	//Validation
	assert.Nil(t, err)
	assert.NotNil(t, response)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

	var apiErr errors.ApiError
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "no country with id AR", apiErr.Message)

}

func TestGetCountriesNoError(t *testing.T) {
	//TODO: Implement me
	//Init
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "AR","name": "Argentina","locale": "es_AR","currency_id": "ARS","decimal_separator": ",","thousands_separator": ".","time_zone": "GMT-03:00","geo_information": {},"states": []}`,
	})

	//Execution
	response, err := http.Get("http://localhost:3000/locations/countries/AR")

	//Validation
	assert.Nil(t, err)
	assert.NotNil(t, response)

	bytes, _ := ioutil.ReadAll(response.Body)

	var country locations.Country
	err = json.Unmarshal(bytes, &country)

	assert.Nil(t, err)
	assert.NotNil(t, country)

	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}
