package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 设置一个 3 秒后的截止时间
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // 确保资源释放

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("任务完成")
	case <-ctx.Done():
		fmt.Println("任务超时:", ctx.Err())
	}

	a := 32              // 32 的二进制表示是 100000
	b := 987628636454310 // 一个很大的数
	result := a & b      // 对两个数进行按位与操作
	fmt.Printf("32 & 987654321 的结果是: %d\n", result)

}
