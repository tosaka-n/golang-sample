### 前回の続き
- socket版チャンク対応

##### チャンク
- HTTP/1.1で分割してリクエストを送るもの
- HTTP/2.0ではHTTP Stream
- Transfer-Encodingヘッダが「chunked」

#### ex
- server.goを起動
- `localhost:18888/` へリクエスト
  - Headerに `Content-Length: 31`
- `localhost:18888/chunked` へリクエスト
  - Headerに`Transfer-Encoding: chunked`

### TCP
- HTTPリクエストでそのまま使える
- ただしレスポンスはHTTPに準拠しない

- [client_chunked_socket.go](./client_chunked_socket.go)
  - HeaderのTransferEncoding: chunkedを確認
  - /chunked を消したエンドポイントへアクセスすると死ぬ 