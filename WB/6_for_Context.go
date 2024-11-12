package main

import (
	"fmt"
	"time"
	"context"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	go work(ctx)

	time.Sleep(11*time.Second)
	fmt.Println("main завершился")
}

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("горутина завершила работу")
			return
		default:
			fmt.Println("горутина работает")
			time.Sleep(time.Second)
		}
	}
}