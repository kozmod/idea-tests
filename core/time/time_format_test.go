package time

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestTimeFormat_1(t *testing.T) {
	fmt.Println(formatTime(time.Now(), "15:04:05"))            // YYYY.MM.DD-hh.mm.ss
	fmt.Println(formatTime(time.Now(), "2006.01.02-15.04.05")) // YYYY.MM.DD-hh.mm.ss
	fmt.Println(formatTime(time.Now(), time.RFC822Z))          // "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	fmt.Println(formatTime(time.Now(), time.RFC850))           // "Monday, 02-Jan-06 15:04:05 MST"
	fmt.Println(formatTime(time.Now(), time.RFC1123Z))         // "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
}

func formatTime(t time.Time, format string) string {
	return t.Format(format)
}

func TestTimeParse_1(t *testing.T) {
	pattern := "15:04:05"
	s := "23:01:00"
	time, err := time.Parse(pattern, s)
	if err != nil {
		panic(err)
	}
	fmt.Println(time)
	assert.Equal(t, time.Second(), 0)
	assert.Equal(t, time.Minute(), 1)
	assert.Equal(t, time.Hour(), 23)
}

func TestTimeParse_AndComparing_1(t *testing.T) {
	pattern := "15:04:05"
	t1, err := time.Parse(pattern, "22:15:05")
	t2, err := time.Parse(pattern, "23:00:00")
	t3, err := time.Parse(pattern, "00:00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(t1.UTC())
	fmt.Println(t1)
	fmt.Println(t1.Sub(t1))
	fmt.Println(t1.Sub(t2), t1.Sub(t2) > 0)
	fmt.Println(t2.Sub(t1), t2.Sub(t1) > 0)
	fmt.Println(t3.Add(t1.Sub(t2)))
	fmt.Println(t3.Add(t2.Sub(t1)))
	fmt.Println(t3.Truncate(t2.Sub(t1)))
}

func TestTimeParse_AndComparing_2(t *testing.T) {
	t.Skip()
	pattern := "15:04:05"
	t1, err := time.Parse(pattern, time.Now().UTC().Format(pattern))
	t2, err := time.Parse(pattern, "23:00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(t1)
	fmt.Println(t2)
	d := t2.Sub(t1)
	fmt.Println(d)
	fmt.Println(t2.Add(d))

	tiker := time.NewTicker(d)
	ctx := context.TODO()
	for {
		select {
		case v := <-tiker.C:
			// The ticker has to be started before f as it can take some time to finish
			tiker = time.NewTicker(d)
			fmt.Println(1, v)
		case v := <-tiker.C:
			fmt.Println(2, v)
		case <-ctx.Done():
			tiker.Stop()
			return
		}
	}
}

func TestParseDuration(t *testing.T) {
	s := "2h"
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, d, 2*time.Hour)

	s = "2m"
	d, err = time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, d, 2*time.Minute)

	s = "2s"
	d, err = time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, d, 2*time.Second)
}

func TestChangeSign(t *testing.T) {
	d, _ := time.ParseDuration("-2m")
	d = -d + (24 * time.Hour)
	fmt.Println(d)
	assert.Equal(t, d, 2*time.Minute+24*time.Hour)
}

func TestGetCurrentTimeUtc_OnlyTimePattern(t *testing.T) {
	pattern := "15:04:05"
	s := "23:01:00"
	parsed, _ := time.Parse(pattern, s)
	fmt.Println(parsed)
	ct := time.Now()
	parsed = parsed.AddDate(ct.Year(), int(ct.Month())-1, ct.Day()-1)
	fmt.Println(parsed)
}

func TestPanicIfNegativeDurationAsTickerArg(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Fail()
		}
	}()
	time.NewTicker(-5 * time.Second)
}

func TestZeroDurationAsTickerArg(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Fail()
		}
	}()
	time.NewTicker(0 * time.Second)
}

func TestTimeParsing_1(t *testing.T) {
	s := "2020-11-03T06:30:00.000Z"
	parsed, err := time.Parse(time.RFC3339, s)
	fmt.Println(parsed)
	fmt.Println(parsed.Format("15:04 PM"))
	fmt.Println(err)
}

func TestMonthsBetween(t *testing.T) {
	pattern := "2006-01-02 15:04:05 -0700 MST"
	from, err := time.Parse(pattern, "2020-04-01 00:00:00 +0000 UTC")
	if err != nil {
		panic(err)
	}
	to, err := time.Parse(pattern, "2021-03-31 00:00:00 +0000 UTC")
	if err != nil {
		panic(err)
	}
	fmt.Println(from, to)
	fmt.Println(to.Sub(from))
	outLayout := "January 2006"
	for to.After(from) {
		fmt.Println(from.Format(outLayout))
		from = from.AddDate(0, 1, 0)
	}
}
