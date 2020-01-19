package main

import (
	"fmt"
	"github.com/magiconair/properties"
)

func main() {
	p := properties.MustLoadFile("conf.properties", properties.UTF8)
	fmt.Println(p.MustGetString("key"))
}
