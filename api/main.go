package main

//structure from: https://medium.com/@godusan/building-a-restful-api-in-go-using-the-gin-framework-a-step-by-step-tutorial-part-1-2-70372ebfa988

import (
	"time"
	"worldmaptools/wmt/configs"
	"worldmaptools/wmt/internal/rest_api/database"
	"worldmaptools/wmt/internal/rest_api/handlers"
	"worldmaptools/wmt/internal/rest_api/repositories"
	"worldmaptools/wmt/internal/rest_api/services"
	serve "worldmaptools/wmt/server"
	routes "worldmaptools/wmt/server/router"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	config := configs.NewConfig()

	client, err := database.NewMongoClient(database.Config{
		DBSource:          config.Database.DatabaseSource,
		MaxOpenConns:      25,
		MaxIdleConns:      25,
		ConnMaxIdleTime:   15 * time.Minute,
		ConnectionTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database client")
		return
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Error().Msgf("Failed to close database client: %v", err)
		}
	}()

	// Initialize repositories
	countryRepo := repositories.NewCountryRepository(client.DB)

	//Initialize services
	countryService := services.NewCountryService(countryRepo)

	// Pass services to handlers
	userHandler := handlers.NewCountryHandler(countryService)

	cors := config.CorsNew()

	router := gin.Default()
	router.Use(cors)

	routes.RegisterPublicEndpoints(router, userHandler)

	server := serve.NewServer(log.Logger, router, config)
	server.Serve()
}
