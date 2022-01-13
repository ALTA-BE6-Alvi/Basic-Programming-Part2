package main

import (
	"fmt"
	"strconv"
	"math"
)

func PrimeNumber(number int) bool {
	count := 1

	if number <= 1 {
		return false
	}
	if number == 2 {
		return true
	}
	if number > 2 && number % 2 == 0 {
		return false
	}
	
	var num float64 = float64(number)
    max_div := int(math.Sqrt(num))
	
	i := 3
    for i <= max_div + 1 { 
        if number % i == 0 {
            return false
		}
		i += 2
		count++
	}

	return true
}

func FullPrima(n int) bool {
	// write your code
	isFullPrima := true
	nums_string := strconv.Itoa(n)
	
	// check type of variable
	// fmt.Println(num_string, reflect.TypeOf(num_string))

	for i := 0; i < len(nums_string); i++ {
		// Check string
		// fmt.Println(string(nums_string[i]))

		// Convert to int and check prime number
		num_string := string(nums_string[i])
		num_int, err := strconv.Atoi(num_string)

		if err != nil {
			fmt.Println("There is error covert to integer", err)
		}

		if !PrimeNumber(num_int) {
			isFullPrima = false
		}
	}
	
	return isFullPrima
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
