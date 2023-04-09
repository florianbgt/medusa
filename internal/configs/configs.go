package configs

import (
	"os"
	"strconv"

	"github.com/google/uuid"
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

func getAPIKey() string {
	var key uuid.UUID

	b, err := os.ReadFile("medusa.KEY")
	if os.IsNotExist(err) {
		key = uuid.New()
		err = os.WriteFile("medusa.KEY", []byte(key.String()), 0644) // 0644 is the file permission
		if err != nil {
			panic(err)
		}
	} else if err == nil {
		parsed_key, err := uuid.ParseBytes(b)
		if err != nil {
			panic(err)
		}
		key = parsed_key
	} else {
		panic(err)
	}

	return key.String()
}

type Configs struct {
	PORT             string
	DEBUG            bool
	API_KEY          string
	DEFAULT_PASSWORD string
	ENABLE_CAMERA    bool
	CAMERA_NAME      string
}

func SetupConfigs() *Configs {
	return &Configs{
		PORT:             getEnvString("PORT", "8080"),
		DEBUG:            getEnvInt("DEBUG", 0) == 1,
		API_KEY:          getAPIKey(),
		DEFAULT_PASSWORD: getEnvString("PASSWORD", "Password/123"),
		ENABLE_CAMERA:    getEnvInt("ENABLE_CAMERA", 1) == 1,
		CAMERA_NAME:      getEnvString("CAMERA_NAME", "/dev/video0"),
	}
}
