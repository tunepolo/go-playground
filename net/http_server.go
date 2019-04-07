package main

import (
	"fmt"
	"net/http"
)

// IndexHandler return Golang Hello World message.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Golang HTTP World")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("0.0.0.0:3000", nil)
}
