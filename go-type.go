package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos = runtime.GOOS
	fmt.Println("goos, ", goos)

	path := os.Getenv("PATH")
	fmt.Println("path: ", path)
}
