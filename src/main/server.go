package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	target string
	code   int
)

func main() {
    var exists bool;
	target, exists = os.LookupEnv("TARGET")
	target = strings.TrimRight(target, "/")
	if !exists || len(target) < 1 {
		log.Fatal("TARGET is not set")
	}

    var err error;
	code, err = strconv.Atoi(os.Getenv("CODE"))
	if err != nil {
        code = http.StatusMovedPermanently    
    }
    
    log.Printf("config> target=%v; code=%v;", target, code)

	if err := http.ListenAndServe(":9000", http.HandlerFunc(redirect)); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	url := target + r.RequestURI
	http.Redirect(w, r, url, code)
}
