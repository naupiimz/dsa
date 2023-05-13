// Find Roots of a Quadratic Equation

// Step 1: Start
// Step 2: Declare variables a, b, c, D, x1, x2, rp and ip;
// Step 3: Calculate discriminant
//          D ← b2-4ac
// Step 4: If D ≥ 0
//               r1 ← (-b+√D)/2a
//               r2 ← (-b-√D)/2a
//               Display r1 and r2 as roots.
//         Else
//               Calculate real part and imaginary part
//               rp ← -b/2a
//               ip ← √(-D)/2a
//               Display rp+j(ip) and rp-j(ip) as roots
// Step 5: Stop

package main

import (
	"fmt"
	"math"
)

func calcDisriminant(a, b, c int) int {
	return (b * 2) - (4 * a * c)
}

func main() {
	a, b, c := 1, 2, 3

	D := calcDisriminant(a, b, c)

	if D >= 0 {
		r1 := (-b + int(math.Sqrt(float64(D)))) / (2 * a)
		r2 := (-b - int(math.Sqrt(float64(D)))) / (2 * a)

		fmt.Println("r1 :", r1, "r2 :", r2)
	} else {
		rp := -b / (2 * a)
		ip := math.Sqrt(-float64(D) / (2 * float64(a)))

		fmt.Println("rp :", rp, "ip :", ip)
	}
}
