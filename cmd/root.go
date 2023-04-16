package cmd

import (
	"flag"
	"fmt"
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
	)

	flag.IntVar(&targetRPM, "rpm", 100, "target RPM (Requests per minute)")
	flag.StringVar(&path, "csv", "", "Path to .CSV file")
	flag.DurationVar(&runDuration, "duration", 0*time.Minute, "Duration of the load test")
	flag.Parse()

	if path == "" {
		fmt.Println("Please provide an csv using -csv flag.")
		return
	}

	if targetRPM == 0 {
		fmt.Println("Please provide a target rpm address using -rpm flag.")
	}

	if runDuration == 0 {
		fmt.Println("Please provide a run duration in minutes using -csv flag.")
		return
	}

	data, err := csv.ReadCSVFile(path, ';')
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	endTime := time.Now().Add(runDuration)

	fmt.Println("======================================================")
	fmt.Println("Starting load-test at:", time.Now().Format("2006-01-02 15:04:05.000000"))
	fmt.Println("======================================================")

	loadtest.Run(data, targetRPM, endTime)

	fmt.Println("======================================================")
	fmt.Println("Done load-testing:", time.Now().Format("2006-01-02 15:04:05.000000"))
	fmt.Println("======================================================")
}
