package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DISCORD_KEY    string
	GEMINI_API_KEY string
	OPEN_AI        string
	role           string = "Hello"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	DISCORD_KEY = os.Getenv("DISCORD_KEY")
	GEMINI_API_KEY = os.Getenv("GEMINI_API_KEY")
}
