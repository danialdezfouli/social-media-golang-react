package config

import (
	"github.com/joho/godotenv"
	"log"
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
	envPath := filepath.Join(filepath.Dir(filename), "../.env")

	var err error
	env, err = godotenv.Read(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Configs = &Config{
		DB:   newDBConfig(env),
		Auth: newAuthConfig(),
		App:  newAppConfig(env),
	}

	return Configs
}

func lookup(key string, fallback string) string {
	if value, ok := env[key]; ok {
		return value
	}
	return fallback
}

func GetString(key string, defaultValue string) string {
	return lookup(key, defaultValue)
}

func GetDuration(key string, fallback time.Duration) time.Duration {
	value := lookup(key, "")
	if value, err := time.ParseDuration(value); err == nil {
		return value
	}
	return fallback
}
