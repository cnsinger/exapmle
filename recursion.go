package main

import "fmt"

func print(n int) {
	fmt.Printf("%d", n)
	if n != 1 {
		fmt.Print("*")
	} else {
		fmt.Print("=")
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	print(n)
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
