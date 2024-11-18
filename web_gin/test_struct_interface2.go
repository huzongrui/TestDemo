package main

import (
	"fmt"
	"unsafe"
)

// 定义一个接口，包含一个获取字段偏移量的方法
type OffsetGetter interface {
	OffsetOfField(fieldName string) uintptr
}

// 定义结构体 ExampleD
type ExampleD struct {
	E int8
	C int32
	A int8
	D int64
	B int16
}

// 实现 OffsetGetter 接口
func (eb ExampleD) OffsetOfField(fieldName string) uintptr {
	fieldOffsets := map[string]uintptr{
		"A": unsafe.Offsetof(eb.A),
		"B": unsafe.Offsetof(eb.B),
		"C": unsafe.Offsetof(eb.C),
		"D": unsafe.Offsetof(eb.D),
		"E": unsafe.Offsetof(eb.E),
	}

	// 使用 range 遍历 map
	for field, offset := range fieldOffsets {
		if field == fieldName {
			return offset
		}
	}

	return uintptr(0) // 字段名无效时返回 0
}

// 打印结构体信息的函数
func PrintFieldOffset(g OffsetGetter, fieldName string) {
	fmt.Printf("Offset of %s: %d bytes\n", fieldName, g.OffsetOfField(fieldName))
}

func main() {
	e := ExampleD{A: 1, B: 2, C: 3, D: 4, E: 0}

	PrintFieldOffset(e, "A")
	PrintFieldOffset(e, "B")
	PrintFieldOffset(e, "C")
	PrintFieldOffset(e, "D")
	PrintFieldOffset(e, "E")
}
