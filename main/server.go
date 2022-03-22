package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}

func handler(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()

	select {
	case <-time.After(2 * time.Second):
		w.Write([]byte("request processed"))
	case <-ctx.Done():
		fmt.Println("request cancelled")
		return
	}
}
