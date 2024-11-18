package main

import (
	"fmt"
)

func main() {
	// 创建一个空的 map，键是字符串类型，值是整数类型
	myMap := make(map[string]int)

	// 添加键值对到 map 中
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["cherry"] = 15

	// 访问 map 中的值
	fmt.Println("Apple:", myMap["apple"])   // 输出: Apple: 5
	fmt.Println("Banana:", myMap["banana"]) // 输出: Banana: 10

	// 检查键是否存在
	value, exists := myMap["cherry"]
	if exists {
		fmt.Println("Cherry exists with value:", value) // 输出: Cherry exists with value: 15
	} else {
		fmt.Println("Cherry does not exist")
	}

	// 删除键值对
	delete(myMap, "banana")

	// 检查删除后的键
	if _, exists := myMap["banana"]; !exists {
		fmt.Println("Banana has been deleted") // 输出: Banana has been deleted
	}

	// 遍历 map
	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
