package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建一个包含 5 个元素的环形链表
	r := ring.New(5)

	// 初始化环形链表中的值
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	// 遍历环形链表中的元素
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

	// 将环形链表中的元素加 10
	for i := 0; i < r.Len(); i++ {
		r.Value = r.Value.(int) + 10
		r = r.Next()
	}

	// 遍历并打印更新后的环形链表
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
}
