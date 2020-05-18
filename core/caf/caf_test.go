package caf

import (
	"fmt"
	"github.com/rakyll/audio/caf"
	"os"
	"testing"
)

func Test(t *testing.T) {

	//file, _ := os.Open("/Users/kozmod/GolandProjects/idea-tests/core/caf/call_waiting_tone_ansi.caf")
	//file, _ := os.Open("/Users/kozmod/GolandProjects/idea-tests/core/caf/ringback_tone_ansi.caf")
	//file, _ := os.Open("/Users/kozmod/GolandProjects/idea-tests/core/caf/sms-received2.caf")
	file, _ := os.Open("/Users/a18388871/GoProjects/idea-tests/core/caf/bass.caf")
	d := caf.New(file)

	err := d.Parse()
	fmt.Println(err)

	fmt.Println(d.SampleRate)
	fmt.Println(d.BitsPerChannel)
	fmt.Println(d.ChannelsPerFrame) //stereo
	fmt.Println(d.AudioDataSize)

	sr := d.SampleRate
	cpf := float64(d.ChannelsPerFrame)
	bch := float64(d.BitsPerChannel)
	bch = 1

	r := float64(d.AudioDataSize) / (sr * cpf * bch)

	fmt.Println(r * 8)
}
