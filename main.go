package main

import (
	"fmt"
	"os"
)

type Page struct {
	url   string
	count int
}

func main() {
	args := os.Args[1:]
	// Check if arguments are correct
	_, argSlice, err := checkAndCastArguments(args)
	if err != nil {
		os.Exit(1)
	}

	// Configure the required files so we can use crawlPage
	cfg, err := Configure(args[0], argSlice[0], argSlice[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// crawlPage
	cfg.wg.Add(1)
	go cfg.crawlPage(args[0])
	cfg.wg.Wait()

	// Print to console
	prettyPrintMap(cfg.pages, args[0])
}
