package main

import (
	"context"
	"fmt"
)

func main() {
	var i int = 42
	var s string = "hello"
	var b bool = true
	var f float64 = 3.14
	var arr []int = []int{1, 2, 3}

	fmt.Println("Hello, Go practice!")
	fmt.Printf("int: %d\n", i)
	fmt.Printf("string: %s\n", s)
	fmt.Printf("bool: %t\n", b)
	fmt.Printf("float64: %.2f\n", f)
	fmt.Printf("slice: %v\n", arr)

	context.TODO()
}
