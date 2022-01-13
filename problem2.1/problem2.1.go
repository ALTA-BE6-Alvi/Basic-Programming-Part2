package main

import "fmt"
import "strconv"

func FaktorBilangan(n int) string {
	// your code here
	var faktorBil string

	for i := 1; i <= n; i++ {
		if n % i == 0 {
			faktorBil += strconv.Itoa(i)
			faktorBil += "\n"
		}
	}

	return faktorBil
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilangan(number))
}
