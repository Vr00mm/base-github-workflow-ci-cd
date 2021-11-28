package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func returnError(err error) {
	newErr := fmt.Sprintf("Error: %v", err)
	if err != nil {
		fmt.Println(newErr)
	}
	os.Exit(1)
}
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		returnError(err)
	}

}
