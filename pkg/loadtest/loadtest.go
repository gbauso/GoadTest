package loadtest

import (
	"fmt"
	"sync"
	"time"

	"github.com/TimLangePN/GoadTest/pkg/httprequest"
)

func Run(data [][]string, rpm int, rampUpPeriod time.Duration, endTime time.Time) {
	fmt.Println("======================================================")
	fmt.Println("Starting load-test at:", time.Now().Format("2006-01-02 15:04:05.000000"))
	fmt.Println("======================================================")

	wg := sync.WaitGroup{}
	dataIndex := 0
	dataLen := len(data)

	interval := rampUpPeriod / time.Duration(rpm)
	ticker := time.NewTicker(interval)

	for time.Now().Before(endTime) {
		wg.Add(1)
		go func() {
			defer wg.Done()

			url := data[dataIndex][0]
			apiKey := data[dataIndex][1]
			httprequest.PerformGetRequest(url, apiKey)

			dataIndex = (dataIndex + 1) % dataLen
		}()

		<-ticker.C
	}

	wg.Wait()
	ticker.Stop()

	fmt.Println("======================================================")
	fmt.Println("Done load-testing:", time.Now().Format("2006-01-02 15:04:05.000000"))
	fmt.Println("======================================================")
}
