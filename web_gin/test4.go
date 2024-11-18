package main

import (
	"fmt"
	"unsafe"
)

type Example struct {
	E int8
	C int32
	A int8
	D int64
	B int16
}

type ExampleB struct {
	D int64
	C int32
	B int16
	A int8
	E int8
}

func main() {
	e := Example{A: 1, B: 2, C: 3, D: 4, E: 0}
	fmt.Printf("Size of Example: %d\n", unsafe.Sizeof(e))
	fmt.Printf("Offset of A: %d\n", unsafe.Offsetof(e.A))
	fmt.Printf("Offset of B: %d\n", unsafe.Offsetof(e.B))
	fmt.Printf("Offset of C: %d\n", unsafe.Offsetof(e.C))
	fmt.Printf("Offset of D: %d\n", unsafe.Offsetof(e.D))
	fmt.Printf("unsafe.Alignof A: %d\n", unsafe.Alignof(e.A))
	fmt.Printf("unsafe.Alignof B: %d\n", unsafe.Alignof(e.B))
	fmt.Printf("unsafe.Alignof C: %d\n", unsafe.Alignof(e.C))
	fmt.Printf("unsafe.Alignof D: %d\n", unsafe.Alignof(e.D))
	fmt.Printf("unsafe.Alignof D: %d\n", unsafe.Alignof(e.E))

	fmt.Printf("#######################################\n")

	e2 := ExampleB{A: 1, B: 2, C: 3, D: 4}
	fmt.Printf("Size of Example: %d\n", unsafe.Sizeof(e2))
	fmt.Printf("Offset of A: %d\n", unsafe.Offsetof(e2.A))
	fmt.Printf("Offset of B: %d\n", unsafe.Offsetof(e2.B))
	fmt.Printf("Offset of C: %d\n", unsafe.Offsetof(e2.C))
	fmt.Printf("Offset of D: %d\n", unsafe.Offsetof(e2.D))

}
