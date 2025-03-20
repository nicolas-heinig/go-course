package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age before:", *agePointer)

	changeAge(agePointer)

	fmt.Println("Adult age:", age)
}

func changeAge(age *int) {
	// return *age - 18
	*age = *age - 18
}
