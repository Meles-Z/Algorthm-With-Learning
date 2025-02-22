package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel:=context.WithCancel(context.Background())
	
	go Worker(ctx)
	time.Sleep(2*time.Second)
	cancel()
	time.Sleep(2*time.Second)
}

func Worker (ctx context.Context){
	for{
		select{
		case <-ctx.Done():
			fmt.Println("Work canceled")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500*time.Microsecond)
		}
	}
}