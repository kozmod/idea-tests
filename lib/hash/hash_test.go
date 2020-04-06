package hash

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"testing"
)

type Apps struct {
	App []App
}

type App struct {
	Name      []string
	Namespace string
}

func Test(t *testing.T) {
	appsl := make([]App, 0)
	appsl = append(appsl, App{[]string{"A1", "B1"}, "NAME_1"})
	appsl = append(appsl, App{[]string{"A2", "B2"}, "NAME_2"})
	apps := Apps{appsl}
	b := hash(apps)
	fmt.Println(b)
}

func hash(apps Apps) [16]byte {
	var appsBytes []byte
	for _, item := range apps.App {
		jsonBytes, _ := json.Marshal(item)
		fmt.Println(jsonBytes)
		appsBytes = append(appsBytes, jsonBytes...)
	}
	return md5.Sum(appsBytes)
}
