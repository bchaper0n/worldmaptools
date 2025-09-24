package routes

import (
	"worldmaptools/wmt/internal/rest_api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicEndpoints(router *gin.Engine, countryHandlers *handlers.Country) {
	router.GET("/countries", countryHandlers.GetAllCountries)
}
