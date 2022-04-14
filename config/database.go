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

func newDBConfig(e Env) *DBConfig {
	return &DBConfig{
		Dialect:  e["DB_CONNECTION"],
		Host:     e["DB_HOST"],
		Port:     e["DB_PORT"],
		Username: e["MYSQL_USER"],
		Password: e["MYSQL_ROOT_PASSWORD"],
		Name:     e["MYSQL_DATABASE"],
		Charset:  "utf8",
	}
}
