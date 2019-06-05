package server

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	currentConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "notchproxy_connections_current",
		Help: "The current number of connections accepted by the proxy",
	})
)

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	// Static localhost endpoint to start
	target := "http://localhost:1201"
	url, err := url.Parse(target)
	if err != nil {
		log.Panicf("Error parsing target Url: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ServeHTTP(res, req)
}

// Run starts the server, which listens for TCP connections
func Run(service string) {
	http.Handle("/metrics", promhttp.Handler())

	proxyHandler := http.HandlerFunc(handleRequestAndRedirect)

	proxyChain := promhttp.InstrumentHandlerInFlight(currentConnections, proxyHandler)

	http.Handle("/", proxyChain)
	if err := http.ListenAndServe(service, nil); err != nil {
		log.Panicf("Error proxying requests: %v", err)
	}
}
