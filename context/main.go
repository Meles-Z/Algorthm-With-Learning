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

func Webscrap(ctx context.Context, url string, result chan<- string) {
	client := http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		result <- fmt.Sprintf("Error happened on %s: %s", url, err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		result <- fmt.Sprintf("Error happened on %s: %s", url, err)
		return
	}
	defer res.Body.Close()

	// Check HTTP status code
	if res.StatusCode != http.StatusOK {
		result <- fmt.Sprintf("Failed to fetch %s: Status %d", url, res.StatusCode)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		result <- fmt.Sprintf("Unreadable response body from %s: %s", url, err)
		return
	}

	result <- fmt.Sprintf("Result from %s:\n%s", url, string(body))
}
func main() {
	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://example.com",
	}
	ch := make(chan string, len(urls))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	for _, url := range urls {
		go Webscrap(ctx, url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	close(ch)
}
