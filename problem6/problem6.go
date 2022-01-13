package main

import (
	"fmt"
)

func FullPrima(n int) bool {
	// write your code
}

func main() {
	fmt.Println(FullPrima(2))  // true
	fmt.Println(FullPrima(3))  // true
	fmt.Println(FullPrima(11)) // false
	fmt.Println(FullPrima(13)) // false
	fmt.Println(FullPrima(23)) // true
	fmt.Println(FullPrima(29)) // false
	fmt.Println(FullPrima(37)) // true
	fmt.Println(FullPrima(41)) // false
	fmt.Println(FullPrima(43)) // false
	fmt.Println(FullPrima(53)) // true
}
