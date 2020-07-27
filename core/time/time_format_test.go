package time

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func TestTimeFormat_1(t *testing.T) {
	fmt.Println(formatTime(time.Now(), "2006.01.02-15.04.05")) // YYYY.MM.DD-hh.mm.ss
	fmt.Println(formatTime(time.Now(), time.RFC822Z))          // "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	fmt.Println(formatTime(time.Now(), time.RFC850))           // "Monday, 02-Jan-06 15:04:05 MST"
	fmt.Println(formatTime(time.Now(), time.RFC1123Z))         // "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
}

func formatTime(t time.Time, format string) string {
	return t.Format(format)
}

func TestTimeParse_1(t *testing.T) {
	s := "23:01:00"
	time, err := time.Parse("15:04:05", s)
	if err != nil {
		panic(err)
	}
	fmt.Println(time)
	assert.Equal(t, time.Second(), 0)
	assert.Equal(t, time.Minute(), 1)
	assert.Equal(t, time.Hour(), 23)
}
