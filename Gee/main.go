package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leilei3167/Framework/Gee/gee"
)

func main() {

	e := gee.New()
	e.GET("/", indexHandler)
	e.GET("/hello", indexHandler)

	log.Fatal(e.Run(":9090"))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
