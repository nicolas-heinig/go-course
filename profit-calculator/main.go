package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Profit Calculator")

	revenue, err := readValue("Enter revenue: ")
	expenses, err := readValue("Enter expenses: ")
	taxRate, err := readValue("Enter tax rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	storeResult(earningsBeforeTax, profit, ratio)

	fmt.Println("======")
	fmt.Printf("Earnings before tax: %.1f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f%%\n", ratio)
}

func readValue(prompt string) (float64, error) {
	var value float64
	fmt.Print(prompt)
	fmt.Scan(&value)

	if value < 0 {
		return 0, fmt.Errorf("Invalid input: %.2f. Must be greater than 0", value)
	}

	if value == 0 {
		return 0, fmt.Errorf("Invalid input: %.2f. Must be greater than 0", value)
	}

	return value, nil
}

func storeResult(earningsBeforeTax float64, profit float64, ratio float64) {
	result := fmt.Sprintf("Earnings before tax: %.1f\nProfit: %.1f\nRatio: %.3f%%\n", earningsBeforeTax, profit, ratio)
	os.WriteFile("results.txt", []byte(result), 0644)
}
