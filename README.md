# brew-note-api

Windows 環境ではそのまま docker 立ち上げても air のホットリロードが聞かないので、wsl 上から実行する

```bash
wsl
docker compose up
```

リビルド

```sh
docker-compose build --no-cache api
```

フォーマット

```sh
go fmt ./...
```

docker logs --tail 100 -f brew-note-api
docker-compose restart webserver

## 滅びの呪文

```sh
docker-compose down --rmi all --volumes --remove-orphans
```
