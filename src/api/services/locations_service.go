package services

import (
	"unitesttraining/src/api/domain/locations"
	"unitesttraining/src/api/providers/locations_provider"
	"unitesttraining/src/api/utils/errors"
)

type locationService struct {
}

type locationServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

func NewLocationService() locationServiceInterface {
	return &locationService{}
}

func (ls *locationService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryId)
}
