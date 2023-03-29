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
	API_KEY          string
	PASSWORD         string
	TOKEN_EXPIRATION int // in seconds
	SQLITE_FILE_PATH string
}

func SetupConfigs() *Configs {
	return &Configs{
		API_KEY:          getEnvString("API_KEY", "change-on-prod"),
		PASSWORD:         getEnvString("PASSWORD", "password"),
		TOKEN_EXPIRATION: getEnvInt("TOKEN_EXPIRATION", 60*60),
		SQLITE_FILE_PATH: getEnvString("SQLITE_FILE_PATH", "app.db"),
	}
}
