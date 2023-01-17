package locations_provider

import (
	"net/http"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountryRestError(t *testing.T) {
	rest.StartMockupServer()

	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: 0,
	})

	// Execution:
	country, err := GetCountry("AR")

	// Validation:
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient error when getting country AR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	rest.StartMockupServer()

	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []}`,
	})
	// Execution:
	country, err := GetCountry("AR")

	// Validation:
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	rest.StartMockupServer()
	// Execution:
	country, err := GetCountry("AR")

	// Validation:
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country AR", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	rest.StartMockupServer()
    
    rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 0,"name": "Argentina","locale": "es_AR","currency_id": "ARS","decimal_separator": ",","thousands_separator": ".","time_zone": "GMT-03:00","geo_information": {},"states": []}`,
	})

	// Execution:
	country, err := GetCountry("AR")

	// Validation:
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for AR", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	rest.StartMockupServer()
	// Execution:
	country, err := GetCountry("AR")

	// Validation:
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, 24, len(country.States))
}
