package main

import "fmt"

func main() {
	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type=%T value=%v\n", u8, u8)

	fmt.Println(1+5, 1+4)
	fmt.Println(10 - 1)
	fmt.Println(10 / 2)
	fmt.Println(10 / 3)
	fmt.Println(10.0 / 3)
	fmt.Println(10 / 3.0)
	fmt.Println(10 % 2)

	x := 1
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)
	fmt.Println(1 << 0)
	fmt.Println(1 << 1)
	fmt.Println(1 << 2)
	fmt.Println(1 << 3)
	fmt.Println(1 << 4)
}
