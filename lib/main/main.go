package main

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

type ExecuteJidResp struct {
	Info   []interface{}
	Return []map[string]interface{}
}

var yaml = `info: []
return:
  - foo.com:
      some_value:
        1.1.1.1: true
      some_value2:
        hey: hi`

func TestReadYmlFromString(t *testing.T) {
	var Resp ExecuteJidResp

	viper.SetConfigType("yml") //document type

	if err := viper.ReadConfig(strings.NewReader(yaml)); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Resp); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Resp)
}

func main() {
	testSuite := []testing.InternalTest{
		{
			Name: "TestReadYmlFromString",
			F:    TestReadYmlFromString,
		},
	}
	testing.Main(nil, testSuite, nil, nil)
}
