package checker

import (
	"net/http"
	"time"
)

func CheckAPI(url string, ch chan<- Result) {
	start := time.Now()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	latency := time.Since(start)

	if err != nil {
		ch <- Result{URL: url, Error: err}
		return
	}
	defer resp.Body.Close()

	ch <- Result{
		URL:        url,
		StatusCode: resp.StatusCode,
		Latency:    latency,
	}
}
