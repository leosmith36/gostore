package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	host string
	port int
}

func getEnvs() (c *config) {
	c = new(config)

	if val, exists := os.LookupEnv("APP_ENV"); !exists || val == "development" {
		if err := godotenv.Load("../../.env.local"); err != nil {
			log.Print("failed to load .env.local")
		}
		if err := godotenv.Load("../../.env.development"); err != nil {
			log.Fatal("failed to load .env.development")
		}
	}

	c.host = getStringEnv("SERVER_HOST", "localhost")
	c.port = getIntEnv("SERVER_PORT", 4535)

	return
}

func getStringEnv(key string, def string) (val string) {
	var exists bool
	if val, exists = os.LookupEnv(key); !exists {
		return def
	}

	return
}

func getIntEnv(key string, def int) (val int) {
	var (
		sval   string
		exists bool
		err    error
	)

	if sval, exists = os.LookupEnv(key); !exists {
		return def
	}

	if val, err = strconv.Atoi(sval); err != nil {
		log.Fatalf("invalid env: %s", key)
	}

	return
}
