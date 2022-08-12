package main

//go run hello.go
//go test
//go install

import (
	"fmt"
)

//var a, b int = 3,5

func main() {
	hello := HelloGo()
	fmt.Println(hello)
	// Added comments
	a, b := 3, 5
	fmt.Printf("Hello, new gopher!\n")

	a, b = Swap2(a, b)

	fmt.Printf("%v + %v = %v\n", a, b, Add2(a, b, 3))
	//factorial 5 = 120
	fmt.Println(Fact(5))
}

// Add dd
func Add(n1, n2 int) int {
	return n1 + n2
}

// Add2 1 or more arguments
func Add2(ints ...int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

// Swap ss
func Swap(n1, n2 int) (int, int) {
	return n2, n1
}

// Swap2 ss
func Swap2(n1, n2 int) (x int, y int) {
	x, y = n2, n1
	return
}

// Fact recursive function
func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}

// HelloGo dd
func HelloGo() string {
	return "Welcome to Go Lang"
}

//upto 13
