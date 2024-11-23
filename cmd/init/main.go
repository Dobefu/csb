package init

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("No .env file found. Please copy it from the .env.example and enter your credentials")
	}
}
