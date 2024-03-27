# migration

https://github.com/rubenv/sql-migrate
## docker composeでDB 起動した後のmigration

```shell
# 差分の確認
docker compose exec migrations sql-migrate status -config="dbconfig.yml" -env="development"

# 差分の反映
docker compose exec migrations sql-migrate up -config="dbconfig.yml" -env="development"

# 差分の戻し
docker compose exec migrations sql-migrate down -config="dbconfig.yml" -env="development"

```

## ラズパイでmigration(ローカルにDBサーバー立ち上げた場合)
```sh
# imageの作成
sudo docker build -t migrations .

# 差分の確認
sudo docker run --rm -v .:/migrations -it --add-host host.docker.internal:host-gateway migrations sql-migrate status -config="dbconfig.yml" -env="raspberry"

# 差分の反映
sudo docker run --rm -v .:/migrations -it --add-host host.docker.internal:host-gateway migrations sql-migrate status -config="dbconfig.yml" -env="raspberry"

# 差分の戻し
sudo docker run --rm -v .:/migrations -it --add-host host.docker.internal:host-gateway migrations sql-migrate status -config="dbconfig.yml" -env="raspberry"
```