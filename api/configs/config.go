package configs

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"worldmaptools/wmt/internal/rest_api/constants"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	Server   serverConfig
	Database databaseConfig
}

type serverConfig struct {
	Address string
}

type databaseConfig struct {
	//DatabaseDriver string
	DatabaseSource string
}

func NewConfig() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := &Config{
		Server: serverConfig{
			Address: "localhost:8080",
		},
		Database: databaseConfig{
			DatabaseSource: "mongodb://" + GetEnvOrPanic(constants.EnvKeys.DBUsername) + ":" + GetEnvOrPanic(constants.EnvKeys.DBPassword) + "@" + GetEnvOrPanic(constants.EnvKeys.DBSource),
		},
	}

	return c
}

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %s not set", key))
	}

	return value
}

func (conf *Config) CorsNew() gin.HandlerFunc {
	allowedOrigin := GetEnvOrPanic(constants.EnvKeys.CorsAllowedOrigin)

	return cors.New(cors.Config{
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{constants.Headers.Origin},
		ExposeHeaders:    []string{constants.Headers.ContentLength},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == allowedOrigin
		},
		MaxAge: constants.MaxAge,
	})
}
