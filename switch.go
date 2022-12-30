package main

import "fmt"

func main() {
	var num int = 100

	switch num {
	case 98, 99:
		fmt.Println("It'is equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}

}
