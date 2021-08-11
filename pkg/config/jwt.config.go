package config

import (
	"os"
)

type JwtConfig struct {
	SECRET string 
}

func LoadJwtConfig() JwtConfig{
	return JwtConfig{
		SECRET: os.Getenv("JWT_SECRET"),
	}
}
