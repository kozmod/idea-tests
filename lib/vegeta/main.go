package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 1_000, Per: time.Second}
	duration := 1000 * time.Second
	h := map[string][]string{
		"Host": {"xxx.com"},
	}
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://192.168.64.5:32250/hp",
		Header: h,
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
			metrics.Add(res)
		}
		wg.Done()
	}()

	for {
		time.Sleep(5 * time.Second)
		metrics.Close()
		fmt.Println("--------------------------------------")
		fmt.Printf("errors: %d\n", len(metrics.Errors))
		fmt.Printf("s: %f\n", metrics.Success)
		fmt.Printf("req: %d\n", metrics.Requests)
		fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
		fmt.Printf("Tonal: %s\n", metrics.Latencies.Total)
		marshalledText, _ := json.MarshalIndent(metrics, "", " ")
		fmt.Printf("metrics: %+v\n", string(marshalledText))
		fmt.Println("**************************************")
	}
}
