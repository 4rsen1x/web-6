package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int

func main() {
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodGet {
			fmt.Fprintf(w, "%d", counter)
		} else if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "invalid request", http.StatusBadRequest)
				return
			}
			countStr := r.FormValue("count")
			count, err := strconv.Atoi(countStr)
			if err != nil {
				http.Error(w, "это не число", http.StatusBadRequest)
				return
			}
			counter += count
			fmt.Fprintf(w, "%d", counter)
		} else {
			http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
		}
	})

	// Запускаем сервер на порту 3333
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}