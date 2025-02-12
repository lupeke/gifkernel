package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	var (
		reqsPerSec int           = 4
		dur        time.Duration = 2
	)

	targetHost := flag.String("host", "localhost:8888", "Target host:port (default: localhost:8888)")
	flag.Parse()

	targets := []vegeta.Target{
		{Method: "GET", URL: fmt.Sprintf("http://%s/index.html", *targetHost)},
		{Method: "GET", URL: fmt.Sprintf("http://%s/script.js", *targetHost)},
		{Method: "GET", URL: fmt.Sprintf("http://%s/styles.css", *targetHost)},
	}

	targeter := vegeta.NewStaticTargeter(targets...)
	attacker := vegeta.NewAttacker()

	rate := vegeta.Rate{Freq: reqsPerSec, Per: time.Second}
	duration := dur * time.Second

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Server Test") {
		metrics.Add(res)
	}
	metrics.Close()

	sep := func(length int, char string) string {
		return strings.Repeat(char, length)
	}

	fmt.Println(sep(40, "-"))
	fmt.Printf("TEST RESULTS FOR %s\n", *targetHost)
	fmt.Println(sep(40, "-"))
	fmt.Printf("Requests: %d\n", metrics.Requests)
	fmt.Printf("Success Rate: %.2f%%\n\n", metrics.Success*100)
	fmt.Printf("P50 (50th percentile): %s\n", metrics.Latencies.P50)
	fmt.Printf("P95 (95th percentile): %s\n", metrics.Latencies.P95)
	fmt.Printf("P99 (99th percentile): %s\n", metrics.Latencies.P99)
	fmt.Printf("Mean Latency: %s\n", metrics.Latencies.Mean)
	fmt.Printf("Max Latency: %s\n", metrics.Latencies.Max)
	fmt.Printf("Total Latency: %s\n", metrics.Latencies.Total.String())
	fmt.Println(sep(40, "="))
}
