package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anuchito/golang-example/context/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	log.Println("runing... \nport: 2021")
	log.Fatal(http.ListenAndServe("127.0.0.1:2021", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.PrintlnContext(ctx, "handler started")
	defer log.PrintlnContext(ctx, "handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hi there!!")
	case <-ctx.Done():
		err := ctx.Err()
		log.PrintlnContext(ctx, "ctx Done:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
