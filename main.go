package main

import (
	"encoding/json"
	"fmt"
	"os"

	"cloudpulse/checker"
)

type Config struct {
	Targets []string `json:"targets"`
}

func loadConfig() Config {
	file, err := os.ReadFile("config/targets.json")
	if err != nil {
		panic("Failed to load targets.json")
	}

	var cfg Config
	json.Unmarshal(file, &cfg)
	return cfg
}

func main() {
	cfg := loadConfig()

	results := make(chan checker.Result)

	for _, url := range cfg.Targets {
		go checker.CheckAPI(url, results)
	}

	for i := 0; i < len(cfg.Targets); i++ {
		result := <-results

		if result.Error != nil {
			fmt.Printf("❌ %s | ERROR: %v\n", result.URL, result.Error)
			continue
		}

		fmt.Printf(
			"✅ %s | Status: %d | Latency: %v\n",
			result.URL,
			result.StatusCode,
			result.Latency,
		)
	}
}
