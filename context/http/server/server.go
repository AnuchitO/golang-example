package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:2021", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("handler started")
	defer log.Println("handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hi there!!")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println("ctx Done:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
