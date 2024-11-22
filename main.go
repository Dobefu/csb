package main

import (
	"log"

	"github.com/Dobefu/csb/cmd/database"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("No .env file found. Please copy it from the .env.example and enter your credentials")
	}

	err = database.Connect()

	if err != nil {
		log.Fatalln("Could not connect to the database: " + err.Error())
	}

	err = database.DB.Ping()

	if err != nil {
		log.Fatalln("Could not connect to the database: " + err.Error())
	}
}

func main() {
	// Main.
}
