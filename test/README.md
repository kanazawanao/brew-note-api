# Test

E2E テスト

```
docker compose build --no-cache api
docker compose up
docker compose exec -T -e CGO_ENABLED=0 api go test -p 1 -count=1 -short -cover -v ./test/e2e/...
```
