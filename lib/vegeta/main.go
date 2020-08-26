package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 1_000, Per: time.Second}
	duration := 50 * time.Second
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
	go func() {
		for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
			metrics.Add(res)
		}
	}()

	for {
		time.Sleep(3 * time.Second)
		metrics.Close()
		fmt.Printf("errors: %d\n", len(metrics.Errors))
		fmt.Printf("s: %f\n", metrics.Success)
		fmt.Printf("s: %d\n", metrics.Requests)
		fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
		fmt.Printf("Tonal: %s\n", metrics.Latencies.Total)
		fmt.Printf("metrics: %s\n", metrics)
	}

}
