package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	hobbies := [3]string{"Boardgames", "LARP", "Gaming"}

	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	slice1 := hobbies[0:2]
	// slice2 := hobbies[:2]

	fmt.Println(slice1)

	slice1 = slice1[1:3]

	fmt.Println(slice1)

	goals := []string{"Learn Go", "Get a job"}
	fmt.Println(goals)

	goals[1] = "Get a job in Go"
	goals = append(goals, "Be awesome")

	fmt.Println(goals)

	products := []Product{
		{"1", "Book", 10.99},
		{"2", "Game", 59.99},
	}

	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
