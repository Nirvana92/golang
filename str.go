package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1. 前缀: 判断str 是否是以ho 开头的
	var str = "hello, world"
	var result = strings.HasPrefix(str, "ho")
	fmt.Println("string prefix result: ", result)

	// 2. 后缀
	result = strings.HasSuffix(str, "wd")
	fmt.Println("string suffix result: ", result)

	// 3. 字符串包含
	result = strings.Contains(str, "llo")
	fmt.Println("string contain result: ", result)

	// 4. 判断子字符串的索引位置
	var index = strings.Index(str, "ll")
	fmt.Println("string index result: ", index)

	// 5. 最后的索引
	index = strings.LastIndex(str, "l")
	fmt.Println("string last index result: ", index)

	// 6. 字符串替换
	var newStr = strings.Replace(str, "h", "H", -1)
	fmt.Println("string replace new str: ", newStr)

	// 7. 统计字符串的次数
	var count = strings.Count(str, "l")
	fmt.Println("string count result: ", count)

	// 8. 返回重复字符串
	newStr = strings.Repeat(str, 2)
	fmt.Println("string repeat new str: ", newStr)

	// 9. 修改字符串的大小写
	newStr = strings.ToLower(str)
	fmt.Println("string to lower: ", newStr)
	newStr = strings.ToUpper(str)
	fmt.Println("string to upper: ", newStr)

}
