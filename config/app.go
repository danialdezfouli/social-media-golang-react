package config

type AppConfig struct {
	Url  string
	Host string
	Port string
}

func newAppConfig(e Env) *AppConfig {
	return &AppConfig{
		Url:  e["APP_HOST"] + ":" + e["APP_PORT"],
		Host: e["APP_HOST"],
		Port: e["APP_PORT"],
	}
}
