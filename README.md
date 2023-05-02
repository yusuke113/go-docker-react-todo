# Echo/Go + ReactのTODOアプリ

## 開発環境立ち上げ
```
# start docker
docker-compose up -d

# start app
docker-compose exec go sh
GO_ENV=dev go run .

# run migrate
docker-compose exec go sh
GO_ENV=dev go run migrate/migrate.go
```

