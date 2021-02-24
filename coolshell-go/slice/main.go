package main

import "fmt"

func print(a, b []int) {
	fmt.Printf("a.len=%v, a.cap=%v", len(a), cap(a))
	fmt.Println(a)
	fmt.Printf("b.len=%v, b.cap=%v", len(b), cap(b))
	fmt.Println(b)
}

func main() {
	a := make([]int, 32)
	b := a[1:16]
	fmt.Println("init...")
	print(a, b)
	a[1] = 11
	fmt.Println("after arrange a[1]=11...")
	print(a, b)
	a = append(a, 1)
	fmt.Println("after arrange append...")
	print(a, b)
	a[2] = 42
	fmt.Println("final arrange a[2]=42......")
	print(a, b)
}
