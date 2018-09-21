/*
This problem was asked by Google.
The area of a circle is defined as πr^2. Estimate π to 3 decimal places using a Monte Carlo method.
Hint: The basic equation of a circle is x2 + y2 = r2.

circle of diameter 1 -> r = 0.5
pi * 0.25

shot -> probability inside circle = p
1 * p = size of the circle
pi = 4p
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	shotCountTotal := 0
	shotCountIn := 0
	currentEstimation := 0.0
	for shotCountTotal < 1000000 {
		x := rand.Float64() - 0.5
		y := rand.Float64() - 0.5
		isInCircle := (x*x)+(y*y) < 0.25
		if isInCircle {
			shotCountIn++
		}
		shotCountTotal++
		p := float64(shotCountIn) / float64(shotCountTotal)
		currentEstimation = 4 * p
	}

	fmt.Println("Result:")
	fmt.Println(currentEstimation)
}
