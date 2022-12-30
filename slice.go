package main

import "fmt"

func main() {
	var arr [6]int

	var sliceArr []int = arr[2:4]

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	for i := 0; i < len(sliceArr); i++ {
		fmt.Println(sliceArr[i])
	}
}
