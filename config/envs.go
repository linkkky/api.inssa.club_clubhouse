package config

import (
	_ "github.com/joho/godotenv/autoload"
)

func ternaryMap(value string, defaultValue string) string {
	return map[bool]string{true: value, false: defaultValue}[len(value) != 0]
}

func getEnv(envName string, defaultValue string) string {
	return ternaryMap(envName, defaultValue)
}
