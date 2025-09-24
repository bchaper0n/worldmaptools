package handlers

import (
	"net/http"
	"worldmaptools/wmt/internal/rest_api/services"

	"github.com/gin-gonic/gin"
)

type Country struct {
	countryService *services.Country
}

func NewCountryHandler(countryService *services.Country) *Country {
	return &Country{countryService: countryService}
}

func (h *Country) GetAllCountries(ctx *gin.Context) {
	allCountries, err := h.countryService.GetAllCountries()
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, err)

		return
	}

	ctx.JSON(http.StatusOK, allCountries)
}
