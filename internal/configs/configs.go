package configs

import (
	"os"
	"strconv"
)

func getEnvString(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	} else {
		return fallback
	}
}

func getEnvInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if exists {
		value, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		return value
	} else {
		return fallback
	}
}

type Configs struct {
	PORT             string
	API_KEY          string
	DEFAULT_PASSWORD string
	TOKEN_EXPIRATION int // in seconds
	SQLITE_FILE_PATH string
}

func SetupConfigs() *Configs {
	return &Configs{
		PORT:             getEnvString("PORT", "80"),
		API_KEY:          getEnvString("API_KEY", "change-on-prod"),
		DEFAULT_PASSWORD: getEnvString("PASSWORD", "Password/123"),
		TOKEN_EXPIRATION: getEnvInt("TOKEN_EXPIRATION", 60*60),
		SQLITE_FILE_PATH: getEnvString("SQLITE_FILE_PATH", "medusa.db"),
	}
}
