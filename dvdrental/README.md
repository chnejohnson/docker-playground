### 介紹

學習 postgres 可以參考學習的 DVD 出租資料庫！
參閱 https://www.postgresqltutorial.com/postgresql-sample-database/

### how to use

建立 image

```
docker build -t dvdrental .
```

建立 container

```
docker run -p 5433:5432 -d --rm --name pgdvd dvdrental
```

- 使用 pgadmin 連到 port 5433，Host 使用 host.docker.internal
- 注意：Dockerfile 的 postgres 版本，最新版無法 restore dvdrental.tar
