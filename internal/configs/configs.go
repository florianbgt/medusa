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
	ENABLE_CAMERA    bool
	CAMERA_NAME      string
}

func SetupConfigs() *Configs {
	return &Configs{
		PORT:             getEnvString("PORT", "80"),
		API_KEY:          getEnvString("API_KEY", "change-on-prod"),
		DEFAULT_PASSWORD: getEnvString("PASSWORD", "Password/123"),
		ENABLE_CAMERA:    getEnvInt("ENABLE_CAMERA", 1) == 1,
		CAMERA_NAME:      getEnvString("CAMERA_NAME", "/dev/video0"),
	}
}
