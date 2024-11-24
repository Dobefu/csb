package init

import (
	"github.com/Dobefu/csb/cmd/logger"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		logger.Fatal("No .env file found. Please copy it from the .env.example and enter your credentials")
	}
}
