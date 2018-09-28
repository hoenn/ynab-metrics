package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var addr = flag.String("listen-address", ":8080", "The address to lsiten on for HTTP requests")

func main() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
