### 画像をPostする
- 画像をPOSTメソッドでContent-Typeの指定をする
- 前回書いたやつ → [client_with_header](../2022-01-13/client_with_header.go)
- Multipartを用いたやり方 → [client_with_header](./client_with_header.go)

### タイムアウト
- curl
  - -m, –max-time	TIME 最大転送時間をTIME秒に制限する
- context.With Timeout()を用いる
  - [client_timeout](./client_timeout.go)