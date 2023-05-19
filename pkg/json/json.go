package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"time"

	. "github.com/TimLangePN/GoadTest/common"
	"github.com/go-playground/validator/v10"
)

func ReadJsonConfig(jsonPath string) (*Config, error) {
	var config *Config

	// Get the current source file directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return &Config{}, nil
	}
	fileDir := filepath.Dir(filename)

	// If the jsonPath is empty, use the default path
	if jsonPath == "" {
		jsonPath = filepath.Join(fileDir, "..", "..", "config.json")
	}

	// Read the config file
	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		fmt.Println("Error:", err)
		return &Config{}, nil
	}

	// Unmarshal the JSON data into the Config struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error:", err)
		return &Config{}, nil
	}

	// Parse the duration string into a time.Duration value
	parsedDuration, err := time.ParseDuration(config.DurationStr)
	if err != nil {
		fmt.Println("Error:", err)
		return &Config{}, nil
	}
	config.Duration = parsedDuration

	err = validateJsonConfig(config)
	if err != nil {
		fmt.Println("Error:", err)
		return &Config{}, nil
	}

	return config, nil
}

func validateJsonConfig(config *Config) error {
	validate := validator.New()

	err := validate.Struct(config)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return errors.New(fmt.Sprintf("Field validation for %s failed on the %s tag in config.json", e.Field(), e.Tag()))
		}
	}

	return nil
}
