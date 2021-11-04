### Commands
- ファイル送信
``` bash
go run file.go
```
- NewReaderで文字列の送信
```
go run reader.go
```
- バイナリファイルの送信
```
go run multi-file.go
```

### NOTE
- vimでの16進数変換 `:%!xxd` はファイル変換されてしまう
- `:wq` で閉じてしまうと16進数表示のままファイル保存されてしまうので注意する
- 16進数 → 10進数 `:%!xxd -r`
