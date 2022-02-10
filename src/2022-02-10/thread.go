package main

import (
	"fmt"
	"time"
)

func sub() {
	count := 0
	for {
		count++
		fmt.Printf("sub: %v\n", count)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	// sub() // ここで無限ループに入るので、以下の for 文は実行されない。
	go sub() // こっちにすると並行処理になる。

	count := 0
	for {
		count++
		fmt.Printf("Main: %v\n", count)
		time.Sleep(5000 * time.Millisecond)
	}
}
