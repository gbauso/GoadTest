package cmd

import (
	"flag"
	"fmt"
	"github.com/TimLangePN/GoadTest/pkg/csv"
	"github.com/TimLangePN/GoadTest/pkg/httprequest"
	"log"
)

func Execute() {
	pFlag := flag.String("p", "", "path for CSV input")
	flag.Parse()

	if *pFlag == "" {
		fmt.Println("Please provide an address using -a flag.")
		return
	}

	data, err := csv.ReadCSVFile(*pFlag, ';')
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	for _, record := range data {
		url := record[0]
		apiKey := record[1]
		httprequest.PerformGetRequest(url, apiKey)
	}
}
