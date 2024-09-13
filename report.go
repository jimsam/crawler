package main

import (
	"fmt"
)

func prettyPrintMap(hashMap map[string]int, website string) {
	fmt.Printf(`
=============================
  REPORT for %v 
=============================
`, website)
	for _, page := range SortMapByIntVal(hashMap, false) {
		fmt.Println("Found ", page.count, " internal links to ", page.url)
	}
}
