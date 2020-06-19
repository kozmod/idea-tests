package utils

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

func AsString(urlFunc map[string]func(http.ResponseWriter, *http.Request)) string {
	var sb strings.Builder
	for k, v := range urlFunc {
		sb.WriteString(k)
		sb.WriteString(" : x")
		sb.WriteString(runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name())
	}
	return sb.String()
}
