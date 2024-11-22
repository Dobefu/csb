package main

import "github.com/joho/godotenv"

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("No .env file found. Please copy it from the .env.example and enter your credentials.")
	}
}

func main() {
	// Main.
}
