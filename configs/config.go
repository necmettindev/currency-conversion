package configs

import (
	"os"
	"strconv"
)

type Config struct {
	Env           string         `env:"ENV"`
	Pepper        string         `env:"PEPPER"`
	Postgres      PostgresConfig `json:"postgres"`
	Redis         RedisConfig    `json:"redis"`
	JWTSecret     string         `env:"JWT_SECRET_KEY"`
	Host          string         `env:"APP_HOST"`
	Port          string         `env:"APP_PORT"`
	FeePercentage float64        `env:"FEE_PERCENTAGE"`
}

func GetConfig() Config {
	perc, err := strconv.ParseFloat(os.Getenv("FEE_PERCENTAGE"), 64)

	if err != nil {
		perc = 0.01
	}

	return Config{
		Env:           os.Getenv("ENV"),
		Pepper:        os.Getenv("PEPPER"),
		Postgres:      GetPostgresConfig(),
		JWTSecret:     os.Getenv("JWT_SECRET_KEY"),
		Host:          os.Getenv("APP_HOST"),
		Port:          os.Getenv("APP_PORT"),
		Redis:         GetRedisConfig(),
		FeePercentage: perc,
	}
}
