package main

import (
	"fmt"
	"net/http"
)

// computeE computes the approximation of e by running a fixed number of iterations.
//
//go:noinline
func computeE(iterations int64) float64 {
	res := 2.0
	fact := 1.0

	for i := int64(2); i < iterations; i++ {
		fact *= float64(i)
		res += 1 / fact
	}
	return res
}

func main() {
	addr := ":9096"
	http.HandleFunc("/hello", HelloServer)

	fmt.Printf("Starting server on: %+v\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Failed to run http server: %v\n", err)
	}
}

//go:noinline
func HelloServer(w http.ResponseWriter, r *http.Request) {
	traceId := r.Header.Get("traceId")
	w.Header().Set("traceId", traceId)
	fmt.Fprintf(w, "hello world test", r.URL.Path[1:])
	
}
