## 前処理
- server
  - server.crt
  - server.go
  - server.key
- client
  - ca.crt
  - ca.key
  - openssl.cnf
## 実行
- curl https://localhost:18443
  client
  ```
  curl: (60) SSL certificate problem: unable to get local issuer certificate
  ```
  server
  ```server
  2022/04/21 07:32:52 http: TLS handshake error from 172.19.0.1:58864: remote error: tls: unknown certificate authority
  ```
- ブラウザアクセス
  - **この接続ではプライバシーが保護されません**

## クライアント証明書作成
- openssl genrsa -out client.key 2048

→ client.keyが生成される

- openssl req -new -nodes -sha256 -key client.key -out client.csr -config openssl.cnf

→ client.csrが生成される

- openssl x509 -req -days 365 -in client.csr -sha256 -out client.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions Client

→ client.crt が生成される

