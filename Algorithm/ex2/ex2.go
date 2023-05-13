// Find the largest number between a, b, c
package main

import "fmt"

func main() {
	a, b, c := 20, 5, 100

	if a > b {
		if a > c {
			fmt.Println("A is the largest number")
		} else {
		}
		fmt.Println("C is the largest number")
	} else {
		if b > c {
			fmt.Println("B is the largest number")
		} else {
			fmt.Println("C is the largest number")
		}
	}
}
