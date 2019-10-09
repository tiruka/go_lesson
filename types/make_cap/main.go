package main

import "fmt"

func main() {
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0, 9)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)
	b := make([]int, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	b = append(b, 1, 2, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)

	// c := make([]int, 5)
	c := make([]int, 0)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}
