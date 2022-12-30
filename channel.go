package main

import (
	"fmt"
	"time"
)

// 协程之间的通信通过 chan 来共享数据
func main() {
	var ch chan string

	ch = make(chan string)

	// ch <- str : 表示用通过ch 发送变量 str
	// str = <- ch : 表示变量str 从通道ch 接受数据

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Nirvana"
	ch <- "hangzhou"
	ch <- "beijing"
	ch <- "london"
}

func getData(ch chan string) {
	var str string

	for {
		str = <-ch
		fmt.Println("str: ", str)
	}
}
