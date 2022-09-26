package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}
}

func AssertEnv(key string) string {
	v := os.Getenv(key)
	if key == "" {
		logrus.Fatalf("missing environment variable: %s", key)
	}
	return v
}
