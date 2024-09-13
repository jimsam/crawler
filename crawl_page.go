package main

import (
	"errors"
	"fmt"
)

func (cfg *config) crawlPage(rawURL string) error {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	cfg.mu.RLock()
	if len(cfg.pages) >= cfg.maxPages {
		return errors.New("Max pages reached! process stoped before finish crawling!")
	}
	cfg.mu.RUnlock()
	normalizedRawURL, err := NormalizeURL(rawURL)
	if err != nil {
		return err
	}
	normalizedStartingURL, err := NormalizeURL(cfg.baseURL)
	if err != nil {
		return err
	}
	if normalizedRawURL != normalizedStartingURL {
		return fmt.Errorf("Found a link to different website: %s", normalizedRawURL)
	}
	cfg.mu.RLock()
	_, ok := cfg.pages[rawURL]
	cfg.mu.RUnlock()
	if !ok {
		cfg.mu.Lock()
		cfg.pages[rawURL] = 1
		cfg.mu.Unlock()
		pageHTML, err := getHTML(rawURL)
		if err != nil {
			return err
		}
		urls, err := getURLsFromHTML(pageHTML, cfg.baseURL)
		if err != nil {
			return err
		}
		for _, url := range urls {
			cfg.wg.Add(1)
			go cfg.crawlPage(url)
		}
	} else {
		cfg.mu.Lock()
		cfg.pages[rawURL] += 1
		cfg.mu.Unlock()
		return nil
	}

	return nil
}
