package config

type AppConfig struct {
	Url        string
	Host       string
	Port       string
	Production bool
}

func newAppConfig() *AppConfig {
	return &AppConfig{
		Production: GetBoolean("Production", false),
		Url:        GetEnv("APP_HOST") + ":" + GetEnv("APP_PORT"),
		Host:       GetEnv("APP_HOST"),
		Port:       GetEnv("APP_PORT"),
	}
}
