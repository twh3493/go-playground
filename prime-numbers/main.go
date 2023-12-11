package main

import "fmt"

func main() {
	fmt.Println("HW-1 Prime Numbers.")
	for a := 2; a <= 100; a++ {
		var c bool = true
		for b := 2; b < a-1; b++ {
			if a%b == 0 {
				c = false

			}
		}
		if c == true {
			fmt.Println("", a)

		}
	}
}
