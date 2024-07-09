package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
}

type PostgresConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
}

type ServerConfig struct {
	RESERVATION_PORT string
	PAYMENT_PORT     string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error while loading .env file: %v", err)
	}

	return &Config{
		Postgres: PostgresConfig{
			DB_HOST:     cast.ToString(coalesce("DB_HOST", "localhost")),
			DB_PORT:     cast.ToString(coalesce("DB_PORT", "5432")),
			DB_USER:     cast.ToString(coalesce("DB_USER", "postgres")),
			DB_NAME:     cast.ToString(coalesce("DB_NAME", "reservation_service")),
			DB_PASSWORD: cast.ToString(coalesce("DB_PASSWORD", "password")),
		},
		Server: ServerConfig{
			RESERVATION_PORT: cast.ToString(coalesce("RESERVATION_PORT", ":50052")),
			PAYMENT_PORT:     cast.ToString(coalesce("PAYMENT_PORT", ":50053")),
		},
	}
}

func coalesce(key string, value interface{}) interface{} {
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}
	return value
}
