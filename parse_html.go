package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"io"
	"net/http"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	htmlNode, err := html.Parse(htmlReader)
	if err != nil {
		return []string{}, fmt.Errorf("There was an error parsing the html: %w", err)
	}

	response := []string{}
	var f func(n *html.Node, res *[]string) error
	f = func(n *html.Node, res *[]string) error {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					urlStruct, err := url.Parse(a.Val)
					if err != nil {
						return fmt.Errorf("There has an error parsing an URL: %w", err)
					}
					if urlStruct.Scheme == "" && urlStruct.Host == "" {
						*res = append(*res, rawBaseURL+a.Val)
					} else {
						*res = append(*res, a.Val)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, res)
		}
		return nil
	}
	f(htmlNode, &response)

	return response, nil
}

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("There was an error fetching data from the web: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode >= 400 {
		return "", fmt.Errorf("There was a %d error: %s", res.StatusCode, body)
	}
	if !strings.Contains(res.Header.Get("Content-Type"), "text/html") {
		return "", errors.New("There response was not of the correct type")
	}

	return string(body), nil
}
