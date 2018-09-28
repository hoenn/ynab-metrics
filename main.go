package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":8080", "The address to lsiten on for HTTP requests")
var token = flag.String("token", "", "Your YNAB access token")

func main() {
	flag.Parse()
	if *token == "" {
		panic("Must run with --token=abc... flag")
	}
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
