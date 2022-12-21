package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App         string
	AppVersion  string
	Environment string // development, staging, production
	HTTPPort    string

	Default_Offset string
	Default_Limit  string

	CategoryServiceGrpcHost string
	CategoryServiceGrpcPort string

	LocationServiceGrpcHost string
	LocationServiceGrpcPort string

	BookServiceGrpcHost string
	BookServiceGrpcPort string

	AuthorServiceGrpcHost string
	AuthorServiceGrpcPort string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "libary_management_service"))
	config.AppVersion = cast.ToString(getOrReturnDefaultValue("APP_VERSION", "1.0.0"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7070"))

	config.Default_Offset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.Default_Limit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.CategoryServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("CATEGORY_SERVICE_HOST", "localhost"))
	config.CategoryServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("CATEGORY_SERVICE_PORT", ":7000"))

	config.LocationServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("LOCATION_SERVICE_HOST", "localhost"))
	config.LocationServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("LOCATION_SERVICE_PORT", ":7000"))

	config.BookServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("BOOK_SERVICE_HOST", "localhost"))
	config.BookServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("BOOK_SERVICE_PORT", ":7000"))

	config.AuthorServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("AUTHOR_SERVICE_HOST", "localhost"))
	config.AuthorServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("AUTHOR_SERVICE_PORT", ":7000"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
