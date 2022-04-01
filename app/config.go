package app

import (
	"log"

	"github.com/joho/godotenv"
)

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}
