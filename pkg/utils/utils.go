package utils

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        panic("No .env file found")
    }
}