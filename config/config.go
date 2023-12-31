package config

import (
	"log/slog"
	"os"
)

type GlobalConfig struct {
    ENVIRONMENT string
    SECRET_KEY []byte
    DATABASE_URI string
    ROOT_DIRECTORY string
}

const DEVELOPMENT = "development"
var SECRET_KEY = "secret"

var Config = GlobalConfig {
    ENVIRONMENT: DEVELOPMENT,
    SECRET_KEY: nil,
    DATABASE_URI: "",
    ROOT_DIRECTORY: "",
}

func InitConfig() {
    // Set the different config variables based on environment variables

    var value string
    var ok bool

    if value, ok = os.LookupEnv("ENVIRONMENT"); ok {
        Config.ENVIRONMENT = value
    } 

    if value, ok = os.LookupEnv("DATABASE_URI"); ok {
        Config.DATABASE_URI = value
    } else if Config.ENVIRONMENT != DEVELOPMENT {
        slog.Error("Database URI not set in non-development environment.")
        os.Exit(1)
    }

    if value, ok = os.LookupEnv("ROOT_DIRECTORY"); ok {
        Config.ROOT_DIRECTORY = value
    } else if Config.ENVIRONMENT != DEVELOPMENT {
        slog.Error("Root directory not set in non-development environment.")
        os.Exit(1)
    } else {
        cwd, err := os.Getwd()
        if err != nil {
            slog.Error("Unable to set root directory.")
            os.Exit(1)
        }
        Config.DATABASE_URI = cwd
    }

    if value, ok = os.LookupEnv("SECRET_KEY"); ok {
        Config.SECRET_KEY = []byte(value)
    } else if Config.ENVIRONMENT != DEVELOPMENT {
        slog.Error("Secret key not set in non-development environment.")
        os.Exit(1)
    } else {
        Config.SECRET_KEY = []byte(SECRET_KEY)
    }
}
