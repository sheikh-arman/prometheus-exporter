package pkg

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
)

var counter = 0
var mu sync.Mutex

func test(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	mu.Unlock()
	fmt.Fprintf(w, "Hello, counter: %d", counter)
	totalRequests.WithLabelValues("200").Inc()
}

func Exporter() {
	http.HandleFunc("/hi", test)
	http.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)
	http.ListenAndServe(":8080", nil)
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests.",
	},
	[]string{"hi"},
)

func init() {
	prometheus.MustRegister(totalRequests)
}
