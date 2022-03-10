### オレオレ証明書を利用するには
- curlコマンドの場合
  - `--insecure`
- chromeの場合
  - chrome://flagsからinsecureを検索してAllowinvalid certificates for resources loaded from localhost: Enable
#### Note
- ルート認証局はオレオレ証明書

### Today's
- Googleの証明書の内容
  - `openssl s_client -connect www.google.com:443 -showcerts < /dev/null 2>&1` (`| grep '[si]:'` をつけるとチェーンの概要)
- 暗号化の処理速度
  - [encrypto_test.go](./encrypto_test.go)
  - `go mod init benchmark`
  - `go test -bench .`
  - ```
    BenchmarkRSAEncryption-4           29878             38865 ns/op
    BenchmarkRSADecryption-4            1196           1041229 ns/op
    BenchmarkAESEncryption-4        10090753               135.2 ns/op
    BenchmarkAESDecryption-4         9872588               116.3 ns/op
    ```

- 利用可能な暗号方式の表示
  - `openssl ciphers -v`