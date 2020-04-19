package main

import "fmt"

func main() {
	n := 5
	fmt.Println(fib(n))
}

func fib(n int) int {

	if n <= 1 {
		return n
	} else {
		return fib(n-2) + fib(n-1)
	}

}
