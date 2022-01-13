- Transport: HTTP, HTTPS HTTPプロキシRoundTripperの実装
- Golangのhttpパッケージを見ていくとhttpを理解できる

### 自由なメソッドの送信
- curl -X DELETE http://localhost:18888 と同義
  - GET/POSTは `-X` の指定不要
- `go run client_delete_method.go`
- newRequest(method, url, body)
- DELETEメソッドの送信
  - [client_delete_method.go](./client_delete_method.go)
- bodyを含む送信(プレーンテキスト)
  - [client_with_body.go](./client_with_body.go)
- bodyを含む送信(オブジェクト)
  - [client_with_body_object.go](./client_with_body_object.go)

### Headerの送信
- curl -H 'Content-Type=image/jpeg' -d "@photo.jpg" http://localhost:18888 と同義