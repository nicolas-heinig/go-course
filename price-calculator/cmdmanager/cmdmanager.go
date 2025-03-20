package cmdmanager

import "fmt"

type CmdManager struct{}

func New() CmdManager {
	return CmdManager{}
}

func (cmd CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter prices")

	var prices []string

	for {
		var price string

		fmt.Println("Price: ")
		fmt.Scanln(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CmdManager) WriteResult(data any) error {
	fmt.Println(data)

	return nil
}
