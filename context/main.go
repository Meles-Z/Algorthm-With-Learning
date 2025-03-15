// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ctx, cancel:=context.WithCancel(context.Background())

// 	go Worker(ctx)
// 	time.Sleep(2*time.Second)
// 	cancel()
// 	time.Sleep(2*time.Second)
// }

// func Worker (ctx context.Context){
// 	for{
// 		select{
// 		case <-ctx.Done():
// 			fmt.Println("Work canceled")
// 			return
// 		default:
// 			fmt.Println("Working...")
// 			time.Sleep(500*time.Microsecond)
// 		}
// 	}
// }

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func execute(ctx context.Context, fn func() int) (int, error) {
	ch := make(chan int, 1)
	go func() {
		ch <- fn()
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func Webscrap(ctx context.Context, url string, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure WaitGroup counter is decremented

	// Create an HTTP client with timeout
	client := &http.Client{}

	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		result <- fmt.Sprintf("Error creating request for %s: %s", url, err)
		return
	}

	// Perform the HTTP request
	res, err := client.Do(req)
	if err != nil {
		result <- fmt.Sprintf("Error fetching %s: %s", url, err)
		return
	}
	defer res.Body.Close()

	// Check HTTP response status
	if res.StatusCode != http.StatusOK {
		result <- fmt.Sprintf("Failed to fetch %s: Status %d", url, res.StatusCode)
		return
	}

	// Read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		result <- fmt.Sprintf("Unreadable response body from %s: %s", url, err)
		return
	}

	// Send successful result
	result <- fmt.Sprintf("Success from %s:\n%s", url, string(body))
}
func main() {
	// List of URLs to scrape
	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://example.com",
	}

	// Create a buffered channel with the same capacity as the number of URLs
	ch := make(chan string, len(urls))

	// Create WaitGroup to track goroutines
	var wg sync.WaitGroup

	// Define a timeout context (e.g., 5 seconds)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Launch goroutines for each URL
	for _, url := range urls {
		wg.Add(1)
		go Webscrap(ctx, url, ch, &wg)
	}

	// Wait for all goroutines to finish in a separate goroutine
	go func() {
		wg.Wait()
		close(ch) // Close the channel after all requests are done
	}()

	// Read from the channel until it's closed
	for result := range ch {
		fmt.Println(result)
	}
}
