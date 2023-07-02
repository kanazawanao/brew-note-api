# migration

https://github.com/rubenv/sql-migrate

```shell
# 差分の確認
docker compose exec migrations sql-migrate status -config="migrations/dbconfig.yml" -env="development"

# 差分の反映
docker compose exec migrations sql-migrate up -config="migrations/dbconfig.yml" -env="development"

# 差分の戻し
docker compose exec migrations sql-migrate down -config="migrations/dbconfig.yml" -env="development"

```
