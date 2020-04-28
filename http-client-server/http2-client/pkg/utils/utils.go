package utils

import (
	"math"
	"strconv"
	"time"
)

func AsSeconds(val string) (time.Duration, error) {
	if f, err := strconv.Atoi(val); err != nil {
		return time.Duration(math.MaxInt32) * time.Second, err
	} else {
		return time.Duration(f) * time.Second, nil
	}
}
