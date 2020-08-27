package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/randomfloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})
	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Int63n(100))
	})
	http.ListenAndServe(":8000", newMux)
}
