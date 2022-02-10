## 本日のテーマ
- Goroutineとチャネル

### Goroutine
- 並行処理のやつ

### チャネル
- Goroutine間で利用するキュー方式のデータ構造

```golang
ch0 := make(chan int)

fmt.Printf("cap(ch0) = %v\n", cap(ch0)) // デフォルトでは cap が 0
fmt.Printf("len(ch0) = %v\n", len(ch0)) // 0

ch1 := make(chan int, 3)                // cap 3 のキュー構造
fmt.Printf("cap(ch1) = %v\n", cap(ch1)) // 3
fmt.Printf("len(ch1) = %v\n", len(ch1)) // 0
// cap(ch0) = 0
// len(ch0) = 0
// cap(ch1) = 3
// len(ch1) = 0
```
- capが0なのでdeadlock
```golang
ch0 := make(chan int)
ch0 <- 1
// fatal error: all goroutines are asleep - deadlock!
```

```golang
ch0 := make(chan int)
i := <-ch0
// fatal error: all goroutines are asleep - deadlock!
```
- capacity: 3に対して値を1つ入れる
```golang
ch1 := make(chan int, 3)
ch1 <- 100
fmt.Printf("cap(ch1) = %v\n", cap(ch1))
fmt.Printf("len(ch1) = %v\n", len(ch1))
// cap(ch1) = 3
// len(ch1) = 1
```
- 取り出すとlenは0になる
```golang
ch1 := make(chan int, 3)
ch1 <- 100
i1 := <-ch1
fmt.Printf("len(ch1) = %v\n", len(ch1))
//  len(ch1) = 0
```
- cap - len < 0 の場合にdeadlock
```golang
ch1 := make(chan int, 3)
ch1 <- 100
ch1 <- 200
ch1 <- 300
ch1 <- 400
// fatal error: all goroutines are asleep - deadlock!
```
- cap: 0でも `go 関数` を書くことでエラーにならずに値を渡せる