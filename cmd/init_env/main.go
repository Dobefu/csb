package init_env

import (
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/joho/godotenv"
)

func Main(envPath string) {
	err := godotenv.Load(envPath)

	if err != nil {
		logger.Fatal("No %s file found. Please copy it from the .env.example and enter your credentials", envPath)
	}
}
