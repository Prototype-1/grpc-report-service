package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    GRPCPort   string
}

func LoadConfig() Config {
    if err := godotenv.Load(); err != nil {
        log.Println("Error: .env file not found")
    }

    return Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", ""),
        DBName:     getEnv("DB_NAME", "grpc_reports"),
        GRPCPort:   getEnv("GRPC_PORT", "50050"),
    }
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
