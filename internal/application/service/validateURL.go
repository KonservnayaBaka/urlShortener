package service

import (
	"net/http"
	url2 "net/url"
	"time"
)

func ValidateUrl(url string) bool {
	parseURL, err := url2.Parse(url)
	if err != nil {
		return false
	}

	if parseURL.Scheme == "" || parseURL.Host == "" {
		return false
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
