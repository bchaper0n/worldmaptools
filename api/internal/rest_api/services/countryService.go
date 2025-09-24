package services

import (
	"net/http"
	"worldmaptools/wmt/internal/rest_api/models"
	"worldmaptools/wmt/internal/rest_api/models/dtos"
	"worldmaptools/wmt/internal/rest_api/repositories"
)

type Country struct {
	countryRepo *repositories.Country
}

func NewCountryService(countryRepo *repositories.Country) *Country {
	return &Country{countryRepo: countryRepo}
}

func (c *Country) GetAllCountries() (*dtos.GetAllCountriesResponse, *models.ErrorResponse) {
	response := &dtos.GetAllCountriesResponse{}

	queriedCountries, err := c.countryRepo.GetAllCountries()
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	response.MapCountriesResponse(queriedCountries)

	return response, nil
}
