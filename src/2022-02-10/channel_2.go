package main

import (
	"fmt"
	"time"
)

func receiver(c chan int) {
	for {
		r := <-c // dequeue
		fmt.Println(r)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go receiver(ch1) // 引数のチャネルは他の Goroutine からの送信を待ち受ける
	go receiver(ch2) // 同上

	fmt.Println("created receiver")
	for i := 0; i < 10; i++ {
		ch1 <- i // enqueue → Goroutine の dequeue が起動
		ch2 <- i // 同上
		time.Sleep(50 * time.Millisecond)
	}
}