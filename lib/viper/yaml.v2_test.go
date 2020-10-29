package viper

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"

	y "gopkg.in/yaml.v2"
)

// ExecuteJidResp2 represends response from Salt API on /execute/<jid> endpoint
type ExecuteJidResp2 struct {
	Info   []interface{}            `yaml:"info"`
	Return []map[string]interface{} `yaml:"return,omitempty"`
}

func TestReadConfig3(t *testing.T) {
	yourBody := `info: []
return:
- foo.com:
	some_value:
		1.1.1.1: true
	some_value2:
	    hey: hi`

	rightBody := `info: []
return:
- foo.com:
    some_value:
      1.1.1.1: true
    some_value2:
      hey: hi`

	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic: %+v\n", p)
		}
	}()

	yamlResp := &ExecuteJidResp2{}
	err := y.Unmarshal([]byte(rightBody), yamlResp)
	if err != nil {
		panic(err)
	}
	spew.Dump(yamlResp)

	yourYamlResp := &ExecuteJidResp2{}
	err = y.Unmarshal([]byte(yourBody), yamlResp)
	if err != nil {
		panic(err)
	}
	spew.Dump(yourYamlResp)
}
