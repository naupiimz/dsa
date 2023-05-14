// Step 1: Start
// Step 2: Declare variables n, factorial and i.
// Step 3: Initialize variables
//           factorial ← 1
//           i ← 1
// Step 4: Read value of n
// Step 5: Repeat the steps until i = n
//      5.1: factorial ← factorial*i
//      5.2: i ← i+1
// Step 6: Display factorial
// Step 7: Stop

package main

import "fmt"

func factorialize(n int) int {
	factorial := 1

	for i := 1; i < n; i++ {
		factorial *= i
	}

	return factorial
}

func main() {
	factorize := factorialize(7)

	fmt.Println(factorize)
}
