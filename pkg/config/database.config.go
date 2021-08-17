package config

import "os"

type DatabaseConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host: os.Getenv("DATABASE_HOST"),
		Port: os.Getenv("DATABASE_PORT"),
		User: os.Getenv("DATABASE_USER"),
		Pass: os.Getenv("DATABASE_PASS"),
		Name: os.Getenv("DATABASE_NAME"),
	}
}

func (d DatabaseConfig) GetConnectionString() string {
	return "postgres://" + d.User + ":" + d.Pass + "@" + d.Host + ":" + d.Port + "/" + d.Name
}
