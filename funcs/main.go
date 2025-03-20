package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)

	transformed := transformNumbers(&numbers, double)

	fmt.Println(transformed)

	sum := sumup(1, 2, 3, 4, 5)

	fmt.Println(sum)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func sumup(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
