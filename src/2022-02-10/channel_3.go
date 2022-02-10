package main

import (
	"fmt"
	"time"
)

func receiver(name string, ch chan int) {
	for {
		i, ok := <-ch // チャネルからの受信の際、2つめの戻り値がチャネルのオープン状態を表す（close なら false）
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int)

	go receiver("1.goroutine", ch1)
	go receiver("2.goroutine", ch1)
	go receiver("3.goroutine", ch1)

	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
	time.Sleep(1 * time.Second) // main関数の終了待ち
}