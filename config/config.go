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
const SECRET_KEY = "development-secret-key"

func GetConfig() *GlobalConfig {
    // Set the different config variables based on environment variables
    c := GlobalConfig {
        ENVIRONMENT: DEVELOPMENT,
        SECRET_KEY: nil,
        DATABASE_URI: "",
        ROOT_DIRECTORY: "",
    }

    var value string
    var ok bool

    if value, ok = os.LookupEnv("ENVIRONMENT"); ok {
        c.ENVIRONMENT = value
    } 

    if value, ok = os.LookupEnv("DATABASE_URI"); ok {
        c.DATABASE_URI = value
    } else if c.ENVIRONMENT != DEVELOPMENT {
        slog.Error("Database URI not set in non-development environment.")
        os.Exit(1)
    }

    if value, ok = os.LookupEnv("ROOT_DIRECTORY"); ok {
        c.ROOT_DIRECTORY = value
    } else if c.ENVIRONMENT != DEVELOPMENT {
        slog.Error("Root directory not set in non-development environment.")
        os.Exit(1)
    } else {
        cwd, err := os.Getwd()
        if err != nil {
            slog.Error("Unable to set root directory.")
            os.Exit(1)
        }
        c.DATABASE_URI = cwd
    }

    if value, ok = os.LookupEnv("SECRET_KEY"); ok {
        c.SECRET_KEY = []byte(value)
    } else if c.ENVIRONMENT != DEVELOPMENT {
        slog.Error("Secret key not set in non-development environment.")
        os.Exit(1)
    }

    return &c
}
