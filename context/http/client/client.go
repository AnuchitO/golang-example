package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:2021", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}
