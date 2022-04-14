package config

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
}

func newDBConfig() *DBConfig {
	return &DBConfig{
		Dialect:  GetEnv("DB_CONNECTION"),
		Host:     GetEnv("DB_HOST"),
		Port:     GetEnv("DB_PORT"),
		Username: GetEnv("DB_USER"),
		Password: GetEnv("DB_ROOT_PASSWORD"),
		Name:     GetEnv("DB_DATABASE"),
		Charset:  "utf8",
	}
}
