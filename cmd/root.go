package cmd

import (
	"flag"
	"fmt"
	"github.com/TimLangePN/GoadTest/config"
	"github.com/TimLangePN/GoadTest/pkg/csv"
	"github.com/TimLangePN/GoadTest/pkg/loadtest"
	"log"
	"time"
)

func Execute() {

	var (
		targetRPM   int
		path        string
		runDuration time.Duration
		jsonPath    string
	)

	flag.IntVar(&targetRPM, "rpm", 0, "target RPM (Requests per minute)")
	flag.StringVar(&path, "csv", "", "Path to .CSV file")
	flag.DurationVar(&runDuration, "duration", 0*time.Minute, "Duration of the load test")
	flag.StringVar(&path, "json", "", "Path to json config file")
	flag.Parse()

	config, err := config.GetConfig(jsonPath, path, targetRPM, runDuration)
	if err != nil {
		fmt.Println("Error getting config:", err)
		return
	}

	data, err := csv.ReadCSVFile(config.CSV, ';')
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}
	endTime := time.Now().Add(config.Duration)

	loadtest.Run(data, config.RPM, endTime)

}
