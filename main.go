package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var port int64

func init() {
	flag.Int64Var(&port, "port", 8080, "the port to listen on")
}

func main() {
	flag.Parse()

	// introduce random failures
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(100) < 10 {
		os.Exit(1)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello from port %d", port)))
	})))
}
