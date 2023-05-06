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

func getEnvFloat(key string, fallback float64) float64 {
	value, exists := os.LookupEnv(key)
	if exists {
		value, err := strconv.ParseFloat(value, 64)
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
	PRINTER_BAUD     int
	PRINTER_NAME     string
	PRINTER_MIN_X    float64
	PRINTER_MAX_X    float64
	PRINTER_MIN_Y    float64
	PRINTER_MAX_Y    float64
	PRINTER_MIN_Z    float64
	PRINTER_MAX_Z    float64
}

func SetupConfigs() *Configs {
	return &Configs{
		PORT:             getEnvString("PORT", "8080"),
		DEBUG:            getEnvInt("DEBUG", 0) == 1,
		API_KEY:          getAPIKey(),
		DEFAULT_PASSWORD: getEnvString("PASSWORD", "Password/123"),
		ENABLE_CAMERA:    getEnvInt("ENABLE_CAMERA", 1) == 1,
		CAMERA_NAME:      getEnvString("CAMERA_NAME", "/dev/video0"),
		PRINTER_NAME:     getEnvString("PRINTER_NAME", "/dev/ttyUSB0"),
		PRINTER_BAUD:     getEnvInt("PRINTER_BAUD", 115200),
		PRINTER_MIN_X:    getEnvFloat("PRINTER_MIN_X", 0),
		PRINTER_MAX_X:    getEnvFloat("PRINTER_MAX_X", 0),
		PRINTER_MIN_Y:    getEnvFloat("PRINTER_MIN_Y", 0),
		PRINTER_MAX_Y:    getEnvFloat("PRINTER_MAX_Y", 0),
		PRINTER_MIN_Z:    getEnvFloat("PRINTER_MIN_Z", 0),
		PRINTER_MAX_Z:    getEnvFloat("PRINTER_MAX_Z", 0),
	}
}
