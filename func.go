// go 语言的函数相关的内容

package main

import "fmt"

func main() {
	fmt.Println("exec main method ...")

	fmt.Println("calc sum, ", sum(2, 4))
}

// 参数 int 类型的a 和 b
// 方法的返回值, int 类型的值
func sum(a, b int) int {
	return a + b
}

/*
定义一个函数的最简单的格式:
func functionName(param_list) return_list {}
*/
