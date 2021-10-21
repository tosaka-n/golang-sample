### What
- Go言語でのクライアント作成

### NOTE
- [client](../src/client) 追加
- `docker-compose up` でサーバ起動するようにした

### Commands
```
$ docker-compose exec app bash
$ cd ../client
$ go run request.go
$ go run query.go
$ go run head.go
$ go run post.go
```