package main

import (
	"flag"
	"fmt"
	"time"
)

const (
	defaultSleepTime = 365 * 24 * time.Hour
)

func main() {
	var sleepTime time.Duration
	flag.DurationVar(&sleepTime, "time", defaultSleepTime, "Sleep Time (time.Duration)")
	flag.Parse()
	fmt.Println("Sleep:", sleepTime)
	time.Sleep(sleepTime)
}
