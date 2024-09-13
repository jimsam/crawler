package main

import (
	"fmt"
	"net/url"
)

func NormalizeURL(str string) (string, error) {
	urlStruct, err := url.Parse(str)
	if err != nil {
		return "", fmt.Errorf("There was an error parse the string to a URL with error: %w", err)
	}

	return urlStruct.Host, nil
}
