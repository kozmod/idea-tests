package context

import "time"

func doSth(duration time.Duration) <-chan time.Time {
	return time.After(duration)
}
