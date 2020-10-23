package tests

import (
	"encoding/json"
	"fmt"
	"testing"
)

type MyData struct {
	One int    `json:"one"`
	two string `json:"two"`
}

func TestMarshalAndUnmarshal(t *testing.T) {
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in)
	encoded, _ := json.Marshal(in)

	fmt.Println(string(encoded))

	var out MyData
	err := json.Unmarshal(encoded, &out)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", out)
}
