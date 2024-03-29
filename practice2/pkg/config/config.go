package config

import (
	"github.com/joho/godotenv"
	"os"
)

func init() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	_ = godotenv.Load("./pkg/env/.env." + env + ".local")
	if env != "test" {
		_ = godotenv.Load("./pkg/env/.env.local")
	}
	_ = godotenv.Load("./pkg/env/.env." + env)
	_ = godotenv.Load()
}
