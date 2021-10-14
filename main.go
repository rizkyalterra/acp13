package main

import "fmt"

func main() {
	fmt.Println("Feature A")
	fmt.Println("Hello World")
	fmt.Println("Hello World2")
	// fixing bug
	// 1
	// 2
	// 3
	// 5
	fmt.Println(fibo(50)) // fibo(3) + fibo(2)
}

var i = 0

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}
