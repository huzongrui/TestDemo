package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine结束")
				return
			default:
				fmt.Println("Goroutine运行中")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancel() // 主动取消 context
	time.Sleep(1 * time.Second)
}
