package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv (env string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erorr Load .env")
	}

	return os.Getenv(env)
}
