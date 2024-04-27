# brew-note-api

```sh
cp .env.defaults .env
```

```sh
docker compose stop
docker compose up -d
```

リビルド

```sh
docker-compose build --no-cache api
```

フォーマット

```sh
go fmt ./...
```

```sh
docker logs --tail 100 -f brew-note-api
docker compose restart api
```

## 滅びの呪文

```sh
docker-compose down --rmi all --volumes --remove-orphans
```
