package main

import (
	"crypto/tls"
	"flag"
	"github.com/symptog/jitsi-colibri-exporter/collector"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	LogLevel = flag.String("loglevel", "info", "Log level")

	MetricsAddr = flag.String("metrics.addr", ":9210", "Metrics address")
	MetricsPath = flag.String("metrics.path", "/metrics", "Metrics path")

	ColibriUrl = flag.String("colibri.url", "http://127.0.0.1:8080/colibri/stats", "Colibiri URL")

	httpClient *http.Client

)

func init() {

	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 100,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

}

func main() {
	flag.Parse()

	lvl, _ := log.ParseLevel(*LogLevel)
	log.SetLevel(lvl)

	log.Info("Starting Jitsi Colibri Exporter")

	coll := collector.New(httpClient, *ColibriUrl)
	prometheus.MustRegister(coll)

	http.Handle(*MetricsPath, promhttp.Handler())
	log.Fatal(http.ListenAndServe(*MetricsAddr, nil))

}
