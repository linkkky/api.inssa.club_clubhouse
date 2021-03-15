package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func ternaryMap(value string, defaultValue string) string {
	return map[bool]string{true: value, false: defaultValue}[len(value) != 0]
}

func getEnv(envName string, defaultValue string) string {
	return ternaryMap(os.Getenv(envName), defaultValue)
}

// Envs has values for environment variables and the defaults for them
var Envs = map[string]string{
	"CLUBHOUSE_ACCOUNT_UUID":       getEnv("CLUBHOUSE_ACCOUNT_UUID", "E656C404-FC76-43F1-816D-2C4F15C4931E"),
	"CLUBHOUSE_ACCOUNT_USER_ID":    getEnv("CLUBHOUSE_ACCOUNT_USER_ID", "1234"),
	"CLUBHOUSE_ACCOUNT_AUTH_TOKEN": getEnv("CLUBHOUSE_ACCOUNT_AUTH_TOKEN", "24204e84df89ba376adf588e3a045c326e6023c2"),
}
