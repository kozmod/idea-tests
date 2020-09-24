package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", firstHandler(activity(lastHandler())))
	http.ListenAndServe(":80", nil)
}

func firstHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := map[string]string{
			"created_by": "A",
			"name":       "Name",
		}
		rcopy := r.WithContext(context.WithValue(r.Context(), "result", v))
		next.ServeHTTP(w, rcopy)
	})
}

func activity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := getParams(r.Context())
		fmt.Println("activity result", result)
		next.ServeHTTP(w, r)
	})
}

func lastHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := getParams(r.Context())
		fmt.Println("last result", result)
	})
}

func getParams(ctx context.Context) map[string]string {
	if ctx == nil {
		return nil
	}

	result, ok := ctx.Value("result").(map[string]string)
	if ok {
		return result
	}

	return nil
}
