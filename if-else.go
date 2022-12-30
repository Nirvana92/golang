package main

import "fmt"

func main() {
	condition := true

	if condition {
		fmt.Println("condition is true")
	} else {
		fmt.Println("condition is false")
	}

	// 第一部分是初始化语句, 第二部分是判断语句
	if val := 0; val > 0 {
		fmt.Println("val is greater zero")
	}
}
