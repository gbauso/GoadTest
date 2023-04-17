package config

import (
	"fmt"
	. "github.com/TimLangePN/GoadTest/common"
	"github.com/TimLangePN/GoadTest/pkg/json"
	"github.com/TimLangePN/GoadTest/pkg/validation"
	"time"
)

func GetConfig(jsonPath, flagPath string, flagRPM int, runDuration time.Duration) (Config, error) {
	err := validation.ValidateFlags(flagPath, flagRPM, runDuration)
	if err != nil {
		fmt.Println(err, "trying to get config from config.json")
		config, errJson := json.ReadJsonConfig(jsonPath)
		if errJson != nil {
			return Config{}, fmt.Errorf("Error reading config: %v", err)
		}
		return config, nil
	}

	return Config{
		XAPIKey:     "test", // Add appropriate values from flags or other sources
		BaseURL:     "test",
		Route:       "test",
		Concurrency: 50,
		RPM:         flagRPM,
		CSV:         flagPath,
	}, nil
}
