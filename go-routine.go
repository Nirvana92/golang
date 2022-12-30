package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// log.SetOutput(os.Stdout)
	fmt.Println("exec init method ....")
}

func main() {
	// fmt.Println("Hello, World")
	fmt.Println("cpu-num, ", runtime.NumCPU())

	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutines")

	go func() {
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		defer wg.Done()
	}()

	go func() {
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}

		// 通知主线程, 本线程的工作已经完成
		defer wg.Done()
	}()

	fmt.Println("Waiting To Finish")

	wg.Wait()

	fmt.Println("\nTerminating Program")
}
