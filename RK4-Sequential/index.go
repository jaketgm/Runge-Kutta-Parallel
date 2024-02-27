/******************************************************
* Idea:
*	- Take in an input function e.g. x^2*y
*	- Solve dy/dx f(x)
*	- Take in an input e.g. \delta x=0.1
*	- And an initial condition e.g. y(0)=1
* 	- Formula for RK4:
*
* y_{i+1}=y_i+\frac{\Delta x}{6}[k_1+2k_2+2k_3+k_4]
******************************************************/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, deltaX float64

	fmt.Println("Enter initial condition x: ")
	if _, err := fmt.Scan(&x); err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	fmt.Println("Enter initial condition y: ")
	if _, err := fmt.Scan(&y); err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	fmt.Println("Enter a value for Δx: ")
	if _, err := fmt.Scan(&deltaX); err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	answer := rungeKuttaFunction(x, y, deltaX)
	fmt.Printf("Using Runge-Kutta of order 4 with x=%v, y=%v, and Δx=%v, the function evaluates to: %v\n", x, y, deltaX, answer)
}

func evaluateFunction(x, y float64) float64 {
	return math.Pow(x, 2) * y // The given function dy/dx = x^2 * y
}

func rungeKuttaFunction(x, y, deltaX float64) float64 {
	kOne, kTwo, kThree, kFour := rungeKuttaFourthOrder(x, y, deltaX)
	yNew := y + (deltaX/6)*(kOne+2*kTwo+2*kThree+kFour)
	return yNew
}

func rungeKuttaFourthOrder(x, y, deltaX float64) (float64, float64, float64, float64) {
	kOne := evaluateFunction(x, y)
	kTwo := evaluateFunction(x+(0.5*deltaX), y+(0.5*kOne*deltaX))
	kThree := evaluateFunction(x+(0.5*deltaX), y+(0.5*kTwo*deltaX))
	kFour := evaluateFunction(x+deltaX, y+kThree*deltaX)
	return kOne, kTwo, kThree, kFour
}
