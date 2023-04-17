package loadtest

import (
	"fmt"
	"github.com/TimLangePN/GoadTest/pkg/httprequest"
	"sync"
	"time"
)

func Run(data [][]string, rpm int, endTime time.Time) {

	fmt.Println("======================================================")
	fmt.Println("Starting load-test at:", time.Now().Format("2006-01-02 15:04:05.000000"))
	fmt.Println("======================================================")

	wg := sync.WaitGroup{}
	interval := time.Minute / time.Duration(rpm)
	dataIndex := 0
	dataLen := len(data)

	for {
		if time.Now().After(endTime) {
			break
		}
		wg.Add(1)
		go func() {
			defer wg.Done()

			url := data[dataIndex][0]
			apiKey := data[dataIndex][1]
			httprequest.PerformGetRequest(url, apiKey)

			dataIndex = (dataIndex + 1) % dataLen
		}()
		time.Sleep(interval)
	}
	wg.Wait()

	fmt.Println("======================================================")
	fmt.Println("Done load-testing:", time.Now().Format("2006-01-02 15:04:05.000000"))
	fmt.Println("======================================================")
}
