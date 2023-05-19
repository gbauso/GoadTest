package config

import (
	"fmt"
	"time"

	"github.com/TimLangePN/GoadTest/common"
	"github.com/TimLangePN/GoadTest/pkg/json"
	"github.com/TimLangePN/GoadTest/pkg/validation"
)

func GetConfig(jsonPath, flagPath string, flagRPM int, runDuration time.Duration, rampUpPeriod time.Duration) (*common.Config, error) {
	err := validation.ValidateFlags(flagPath, flagRPM, runDuration)
	if err != nil {
		fmt.Println(err, "trying to get config from config.json")
		config, errJson := json.ReadJsonConfig(jsonPath)
		if errJson != nil {
			return &common.Config{}, fmt.Errorf("Error reading config: %v", err)
		}
		return config, nil
	}

	return &common.Config{
		XAPIKey:      "test", // Add appropriate values from flags or other sources
		BaseURL:      "test",
		Route:        "test",
		Concurrency:  50,
		RPM:          flagRPM,
		CSV:          flagPath,
		Duration:     runDuration,
		RampUpPeriod: rampUpPeriod,
	}, nil
}
