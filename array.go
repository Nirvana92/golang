package main

import "fmt"

func main() {
	// 数组的声明方式
	// var identifier [len]type

	var arr [5]int

	for i := 0; i < 5; i++ {
		arr[i] = i * 2
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
