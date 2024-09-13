package main

import (
	"fmt"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            string
	maxPages           int
	mu                 *sync.RWMutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func Configure(startUrl string, concurrent int, maxPages int) (config, error) {
	if concurrent == 0 {
		return config{}, fmt.Errorf("concurrent needs to be bigger that 0")
	}

	cfg := config{
		pages:              map[string]int{},
		baseURL:            startUrl,
		maxPages:           maxPages,
		mu:                 &sync.RWMutex{},
		concurrencyControl: make(chan struct{}, concurrent),
		wg:                 &sync.WaitGroup{},
	}
	return cfg, nil
}
