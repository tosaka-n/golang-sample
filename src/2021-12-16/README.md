# プロキシの利用
### file
- client_proxy.go
  - プロキシのリクエストを投げる
  - サーバ側対応が入っていないため返答は変わらない
### ファイルアクセス
- curl file://$(pwd)/README.md
- file_request.go
  - go run file_request.go
  - `file_request.go` ファイルを取得

### Try
- ResisterProtocol, NewFileTransport を公式docから調べる
- 今までのソースコードに出てきた関数も調べる
