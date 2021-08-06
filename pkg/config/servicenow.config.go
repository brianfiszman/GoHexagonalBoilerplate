package config

import (
	"os"
)

type ServiceNowConfig struct {
	API_URL string
	USER    string
	PASS    string
}

func LoadServiceNowConfig() ServiceNowConfig {
	return ServiceNowConfig{
		os.Getenv("SERVICE_NOW_API_BASE_URL"),
		os.Getenv("SERVICE_NOW_USER"),
		os.Getenv("SERVICE_NOW_PASS"),
	}
}
