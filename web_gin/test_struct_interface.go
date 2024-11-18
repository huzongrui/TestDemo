package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// // 定义一个接口，包含一个获取字段偏移量的方法
// type OffsetGetter interface {
// 	OffsetOfField(fieldName string) uintptr
// }

// // 定义结构体 Example
// type ExampleC struct {
// 	E int8
// 	C int32
// 	A int8
// 	D int64
// 	B int16
// }

// // 实现 OffsetGetter 接口
// func (e ExampleC) OffsetOfField(fieldName string) uintptr {
// 	switch fieldName {
// 	case "A":
// 		return unsafe.Offsetof(e.A)
// 	case "B":
// 		return unsafe.Offsetof(e.B)
// 	case "C":
// 		return unsafe.Offsetof(e.C)
// 	case "D":
// 		return unsafe.Offsetof(e.D)
// 	case "E":
// 		return unsafe.Offsetof(e.E)
// 	default:
// 		return uintptr(0)
// 	}
// }

// // 定义结构体 ExampleB
// type ExampleD struct {
// 	D int64
// 	C int32
// 	B int16
// 	A int8
// 	E int8
// }

// // 实现 OffsetGetter 接口
// func (eb ExampleD) OffsetOfField(fieldName string) uintptr {
// 	switch fieldName {
// 	case "A":
// 		return unsafe.Offsetof(eb.A)
// 	case "B":
// 		return unsafe.Offsetof(eb.B)
// 	case "C":
// 		return unsafe.Offsetof(eb.C)
// 	case "D":
// 		return unsafe.Offsetof(eb.D)
// 	case "E":
// 		return unsafe.Offsetof(eb.E)
// 	default:
// 		return uintptr(0)
// 	}
// }

// // 函数用于打印结构体字段的偏移量
// func PrintFieldOffset(g OffsetGetter, fieldName string) {
// 	fmt.Printf("Offset of %s: %d bytes\n", fieldName, g.OffsetOfField(fieldName))
// }

// func main() {
// 	e := ExampleC{A: 1, B: 2, C: 3, D: 4, E: 0}
// 	eb := ExampleD{A: 1, B: 2, C: 3, D: 4, E: 0}

// 	PrintFieldOffset(e, "A")
// 	PrintFieldOffset(e, "B")
// 	PrintFieldOffset(e, "C")
// 	PrintFieldOffset(e, "D")
// 	PrintFieldOffset(e, "E")

// 	PrintFieldOffset(eb, "A")
// 	PrintFieldOffset(eb, "B")
// 	PrintFieldOffset(eb, "C")
// 	PrintFieldOffset(eb, "D")
// 	PrintFieldOffset(eb, "E")
// }
////
