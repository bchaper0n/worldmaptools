package constants

import (
	"time"
)

var EnvKeys = envKeys{
	Env:               "ENV",
	ServerAddress:     "SERVER_ADDRESS",
	DBUsername:        "MONGO_ROOT_USERNAME",
	DBPassword:        "MONGO_ROOT_PASSWORD",
	CorsAllowedOrigin: "CORS_ALLOWED_ORIGIN",
	DBSource:          "MONGO_SOURCE",
	//DBDriver:          "DB_DRIVER",
}

var Headers = headers{
	Origin:        "Origin",
	ContentLength: "Content-Length",
}

var MaxAge = 12 * time.Hour

type envKeys struct {
	Env               string
	ServerAddress     string
	DBUsername        string
	DBPassword        string
	CorsAllowedOrigin string
	DBSource          string
	//DBDriver          string
}

type headers struct {
	Origin        string
	ContentLength string
}
