package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	adjustedValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value: ", futureValue)
	fmt.Printf("Future Value: %.2f\n", futureValue)

	// fmt.Println("Future Value (adjusted for inflation): ", adjustedValue)
	fmt.Printf("Future Value (adjusted for inflation): %.2f\n", adjustedValue)
}
