package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestPwd(t *testing.T) {
	cmd := exec.Command("pwd")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
