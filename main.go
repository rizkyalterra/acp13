package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(fibo(50)) // fibo(3) + fibo(2)
}

var i = 0

func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}
