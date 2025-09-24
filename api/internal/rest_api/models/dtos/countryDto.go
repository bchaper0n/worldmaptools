package dtos

import "worldmaptools/wmt/internal/rest_api/entities"

type CountryResponse struct {
	Name         string `json:"name,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Capital      string `json:"capital,omitempty"`
	Continent    string `json:"continent,omitempty"`
	Flag         string `json:"flag,omitempty"`
}

type GetAllCountriesResponse struct {
	Countries []*CountryResponse `json:"countries"`
}

func (r *GetAllCountriesResponse) MapCountriesResponse(countries []*entities.Country) {
	for _, countries := range countries {
		country := &CountryResponse{
			Name:         countries.Name,
			Abbreviation: countries.Abbreviation,
			Capital:      countries.Capital,
			Continent:    countries.Continent,
			Flag:         countries.Flag,
		}
		r.Countries = append(r.Countries, country)
	}
}

func (r *CountryResponse) MapCountryResponse(country *entities.Country) {
	r.Name = country.Name
	r.Abbreviation = country.Abbreviation
	r.Capital = country.Capital
	r.Continent = country.Continent
	r.Flag = country.Flag
}
