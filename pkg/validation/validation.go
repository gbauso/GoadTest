package validation

import (
	"fmt"
	"time"
)

func ValidateFlags(path string, targetRPM int, runDuration time.Duration) error {
	if path == "" {
		return fmt.Errorf("Please provide an csv using -csv flag.")
	}

	if targetRPM == 0 {
		return fmt.Errorf("Please provide a target rpm address using -rpm flag.")
	}

	if runDuration == 0*time.Minute {
		return fmt.Errorf("Please provide a run duration in minutes using -csv flag.")
	}
	return nil
}
