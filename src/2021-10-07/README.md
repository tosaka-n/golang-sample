## command
```
docker-compose exec app bash
```
```
cd min_server/
```
```
go run server.go
```
```
curl --http1.0 -v -d "{\"hello\": \"world\"}" -H "Content-Type: application/json" http://localhost:18888
```
```
curl --http1.0 -v -X PUT -d "{\"hello\": \"world\"}" -H "Content-Type: application/json" http://localhost:18888
```
```
curl --http1.0 -v -X DELETE -d "{\"hello\": \"world\"}" -H "Content-Type: application/json" http://localhost:18888
```
```
curl --http1.0 -v -d @test.json -H "Content-Type: application/json" http://localhost:18888
```
```
curl --http1.0 -v -d title="Head First PHP & MySQL" -d author="Lynn Beighley, Michael Morrison" http://localhost:18888
```
```
curl --http1.0 -v --data-urlencode title="Head First PHP & MySQL" --data-urlencode author="Lynn Beighley, Michael Morrison" http://localhost:18888
```
```
curl --http1.0 -F title="The Art of Community" -F author="Jono Bacon" -F attachment-file=@test.txt http://localhost:18888
```
サーバー再起動
```
go run server_handle.go 
```
```
curl --http1.0 -c cookie.txt http://localhost:18888
```
```
curl --http1.0 -b cookie.txt http://localhost:18888
```