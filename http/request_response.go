package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type HttpClientConfig struct {
	Timeout       time.Duration
	Token         string
	MaxConcurrent int
}

type WebResponse struct {
	URL      string
	Response []byte
	Error    error
}

func main() {
	cfg := HttpClientConfig{
		Timeout:       time.Second * 2,
		Token:         "Bearer Token",
		MaxConcurrent: 5,
	}
	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://example.com",
	}
	responses := FetchURLs(context.Background(), cfg, urls)
	for _, resp := range responses {
		if resp.Error != nil {
			fmt.Printf("Error fetching:%s: %v", resp.URL, resp.Error)
			continue
		}
		fmt.Printf("Response from %s:\n%s\n\n", resp.URL, string(resp.Response))
	}
}

func FetchURLs(ctx context.Context, cfg HttpClientConfig, urls []string) []WebResponse {
	var wg sync.WaitGroup
	result := make(chan WebResponse, len(urls))
	semaphore := make(chan struct{}, cfg.MaxConcurrent)

	client := &http.Client{
		Timeout: cfg.Timeout,
	}

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() {
				<-semaphore
			}()
			reqCtx, cancel := context.WithTimeout(ctx, cfg.Timeout)
			defer cancel()
			req, err := http.NewRequestWithContext(reqCtx, "GET", u, nil)
			if err != nil {
				result <- WebResponse{URL: u, Error: err}
				return
			}
			req.Header.Set("Authorization", cfg.Token)

			resp, err := client.Do(req)
			if err != nil {
				result <- WebResponse{URL: u, Error: err}
				return
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				result <- WebResponse{URL: u, Error: err}
				return
			}
			result <- WebResponse{URL: url, Response: body, Error: err}

		}(url)

	}
	go func() {
		wg.Wait()
		close(result)
	}()

	var responses []WebResponse
	for results := range result {
		responses = append(responses, results)
	}
	return responses

}
