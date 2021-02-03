package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Error    error
	Response *http.Response
	Url      string
}

type ClientWithTimeout struct {
	client http.Client
}

func NewClientWithTimeout(d time.Duration) *ClientWithTimeout {
	return &ClientWithTimeout{
		client: http.Client{Timeout: d},
	}
}

func (c *ClientWithTimeout) DoRequestWithContext(ctx context.Context, url string, wg *sync.WaitGroup, ch chan<- Result) {
	if wg != nil {
		defer wg.Done()
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		ch <- Result{Url: url, Error: err}
		return
	}

	resp, err := c.client.Do(request)
	if err != nil {
		ch <- Result{Url: url, Error: err}
		return
	}

	if resp.StatusCode != http.StatusOK {
		ch <- Result{Url: url, Error: fmt.Errorf("bad status: %d", resp.StatusCode)}
		return
	}

	ch <- Result{Url: url, Response: resp}
}

func main() {
	urls := []string{
		"https://yandex.ru/",
		"https://somesite.come",
		"https://ыпвыпывпыапы",
		"https://5.5.5.5",
	}

	client := NewClientWithTimeout(1 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	resultsCh := make(chan Result, len(urls))
	wg := &sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go client.DoRequestWithContext(ctx, url, wg, resultsCh)
	}

	for i := 0; i < 2; i++ {
		res := <-resultsCh
		if res.Error != nil {
			fmt.Printf("URL: %s; Error: %s\n", res.Url, res.Error.Error())
		} else {
			fmt.Printf("URL: %s; OK\n", res.Url)
		}
	}

	cancel()
	wg.Wait()
}
