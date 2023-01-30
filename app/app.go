package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
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
	addr := ":9090"

	http.HandleFunc("/helloworld", HelloServer)

	fmt.Printf("Starting server on: %+v\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Failed to run http server: %v\n", err)
	}
}

func GetHello() (string, string) {
	traceId := uuid.New()
	req, err := http.NewRequest("GET", "http://localhost:9096/hello", nil)
	req.Header.Set("traceId", traceId.String())
	req.Close = true
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		fmt.Print(k)
		fmt.Print(" : ")
		fmt.Println(v)
	}
	traceResp := resp.Header.Get("traceId")
	fmt.Println("trace id " + traceResp)
	responseBody := ""
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		responseBody = string(bodyBytes)
	}
	return responseBody, traceResp
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	hello, header := GetHello()
	w.Header().Set("traceId", header)
	fmt.Fprintf(w, hello, r.URL.Path[1:])
}
