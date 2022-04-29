package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Env map[string]string

type Config struct {
	DB   *DBConfig
	Auth *AuthConfig
	App  *AppConfig
}

var Configs *Config
var env Env

func GetConfig() *Config {
	if Configs != nil {
		return Configs
	}

	_, filename, _, _ := runtime.Caller(0)
	envPath := filepath.Join(filepath.Dir(filename), "../.env.development")

	if _, err := os.Stat(envPath); !errors.Is(err, os.ErrNotExist) {
		env, _ := godotenv.Read(envPath)
		//if err != nil {
		//log.Fatal("Error loading .env.development file")
		//}

		for key, val := range env {
			if os.Getenv(key) != "" {
				fmt.Println("default Value for "+key+" :", os.Getenv(key))
			}
			os.Setenv(key, val)
		}
	}

	Configs = &Config{
		DB:   newDBConfig(),
		Auth: newAuthConfig(),
		App:  newAppConfig(),
	}

	return Configs
}

func GetString(key string, fallback string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return fallback
}

func GetEnv(key string) string {
	return GetString(key, "")
}

func GetBoolean(key string, fallback bool) bool {
	value := GetString(key, "")

	if value == "" {
		return fallback
	}

	if value == "true" || value == "1" {
		return true
	}

	return false
}

func GetDuration(key string, fallback time.Duration) time.Duration {
	value := GetString(key, "")
	if value, err := time.ParseDuration(value); err == nil {
		return value
	}
	return fallback
}
