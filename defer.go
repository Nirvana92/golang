package main

import "fmt"

func main() {
	// 关键字 defer 允许我们推迟到函数返回之前, defer 蕾丝与Java 语言中的 finally 语句块
	func_1()
}

func func_1() {
	fmt.Printf("in func-1 at start\n")
	// 可以更换执行如下两个语句对比结果
	// defer func_2()
	// func_2()
	fmt.Printf("in func-1 at end\n")
}

func func_2() {
	fmt.Printf("func-2 invoke exec...\n")
}
