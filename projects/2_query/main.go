package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}
		fmt.Fprintf(w, "Hello,%s!", name)
	})

	// Запускаем сервер на порту 9000
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
