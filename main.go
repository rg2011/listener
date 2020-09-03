package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		io.Copy(ioutil.Discard, r.Body)
		r.Body.Close()
	}()
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Printf("BODY: %s\n\n", string(body))
	}
	w.WriteHeader(204)
}

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(handler))
}
