### お品書き
- ルート証明書を作成
- ルート証明書をOSにインストール
- サーバ証明書をルート証明書をつかって作成
- HTTPSサーバにアクセス

#### 証明書作成
```
openssl x509 -in server.crt -text -noout
openssl x509 -in server.csr -days 3650 -req -signkey server.key -sha256 -out server.crt
openssl req -in server.csr -text -noout
openssl req -new -sha256 -key server.key -out server.csr
openssl genrsa -out server.key 2048
```

- OSのコンフィグをコピー
  - `cp /etc/ssl/openssl.cnf ./`

#### ルート認証局の証明書スイートを作成
```
$ openssl genrsa -out ca.key 2048
$ openssl req -new -sha256 -key ca.key -out ca.csr -config ./openssl.cnf
$ openssl x509 -in ca.csr -days 365 -req -signkey ca.key -sha256 -out ca.crt -extfile ./openssl.cnf -extensions CA
```

##### 確認
```
$ openssl rsa -in ca.key -text
$ openssl req -in ca.csr -text
$ openssl x509 -in ca.crt -text
```

#### サーバ証明書の作成
```
$ openssl genrsa -out server.key 2048
$ openssl req -new -nodes -sha256 -key server.key -out server.csr -config ./openssl.cnf
$ openssl x509 -req -days 365 -in server.csr -sha256 -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions Server
```

#### curlでリクエスト
```
$ curl --cacert ./ca.crt https://localhost:18443
```