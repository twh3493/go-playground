package main

import "fmt"

func main() {
	// fmt.Println("HW-1 Prime Numbers.")
	for a := 2; a <= 100; a = a + 1 {
		if isPrime(a) {
			fmt.Println(a)
		}
	}
}
func isPrime(x int) bool {
	for b := 2; b < x-1; b++ {
		if x%b == 0 {
			return false
		}

	}
	return true
}
