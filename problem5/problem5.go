package main

import "fmt"

func Pangkat(base, pangkat int) int {
	// your code here
	if pangkat == 0 {
		return 1
	}
	hasil := base
	for i := 2; i <= pangkat; i++ {
		hasil *= base
	}
	return hasil
}

func main() {
	fmt.Println(Pangkat(2, 3))  // 8
	fmt.Println(Pangkat(5, 3))  // 125
	fmt.Println(Pangkat(10, 2)) // 100
	fmt.Println(Pangkat(2, 5))  // 32
	fmt.Println(Pangkat(7, 3))  // 343
}
