package database

import (
    "os"
)

type Config struct {
    Host     string
    Port     string
    User     string
	Name	 string
    Password string
}

func NewConfigPostgres() *Config {
    return &Config{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        User:     os.Getenv("DB_USER"),
		Name:     os.Getenv("DB_NAME"),
        Password: os.Getenv("DB_PASSWORD"),
    }
}
