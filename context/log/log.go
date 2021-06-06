package log

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"

	"log"
)

const requestIDKey = 19

func Println(v ...interface{}) {
	log.Println(v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func PrintlnContext(ctx context.Context, msg ...interface{}) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Println("could not find request ID in context")
		return
	}

	log.Printf("[%d] %s\n", id, fmt.Sprint(msg...))
}

func Decorate(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)

		fn(w, r.WithContext(ctx))
	}
}
