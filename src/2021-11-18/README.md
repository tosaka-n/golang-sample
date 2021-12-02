touch test.txt
### サーバ起動
```
docker-compose exec app bash
./server > server.log
```
### ファイル送信
```
docker-compose exec app bash 
go run multi-file.go
```
### 改行コード確認
```
vim server.log
:%!xxd
→ 改行コードが0d0aだということがわかる
```

テキストデータは `form-data`
> Content-Disposition: form-data;
画像データは `application/octet-stream`
> Content-Type: application/octet-stream

curlでやると？
```
$ ./server > server.log
$ curl -F "name=Michael Jackson" -F "thumbnail=@photo.png" http://localhost:18888 
$ cat server.log
ser-Agent: curl/7.64.1

--------------------------2187e9eba04f760c
Content-Disposition: form-data; name="name"

Michael Jackson
--------------------------2187e9eba04f760c
Content-Disposition: form-data; name="thumbnail"; filename="photo.png"
Content-Type: image/png

```
### Content-Type指定
ファイルを書き換えて `part.Set` で `Content-Type` を指定する
```
go run multi-file-v2.go
User-Agent: Go-http-client/1.1

--5c9941fe0b1fc83e2bdfa39ae6a092093c68848b2149de1fa43b7b98df86
Content-Disposition: form-data; name="name"

Michael Jackson
--5c9941fe0b1fc83e2bdfa39ae6a092093c68848b2149de1fa43b7b98df86
Content-Disposition: form-data; name="thumbnail"; filename="photo.jpg"
Content-Type: image/jpeg
```

### Cookieの送受信

```
$ go run server_cookie.go > server.log
```
```
go run coockie.go
```