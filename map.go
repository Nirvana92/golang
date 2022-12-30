package main

import "fmt"

func main() {
	// map 的声明方式
	// var map_name map[keytype]valuetype

	// 1. 直接初始化map 的数据结构
	var mapLit map[string]int = map[string]int{"name": 1, "age": 2}

	// 2. 通过make 创建一个map
	var nameMap map[string]string = make(map[string]string)
	nameMap["name"] = "Nirvana"
	nameMap["age"] = "2"

	fmt.Println("map: ", mapLit)
	fmt.Println("nameMap: ", nameMap)
}
