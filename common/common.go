package common

import "time"

type Config struct {
	XAPIKey     string `json:"x-api-key" validate:"required"`
	BaseURL     string `json:"baseurl" validate:"required,url"`
	Route       string `json:"route" validate:"required"`
	Concurrency int    `json:"concurrency" validate:"required,gt=0"`
	RPM         int    `json:"rpm" validate:"required,gt=0"`
	CSV         string `json:"csv-path" validate:"required"`
	DurationStr string `json:"duration" validate:"required"`
	Duration    time.Duration
}
