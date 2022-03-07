# Initial Project Configuration

This configuration is designed on top of the IntelliJ editor, so it must be installed along with version 1.17 of the Golang.
After validating that these tools will be saved, it must be verified if the following tools are installed, in case
If you don't have any installed, you can find a link to install it.

- Postgres: Docker must be installed to install a local postgres database. A
  Once the docker is installed, they must be executed
  the following commands from a terminal:
```
docker pull postgres
```
```
docker run --name ms-remision -p 5009:5432 -e POSTGRES_PASSWORD=81943301Te -d postgres
```
- Migrate: for install migrate follow instruction in this link: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- Make: install execute this steps:
init console with admin and execute this command (install of windows)
```
choco install make
```

# Execute commands defined on Makefile:
- make postgrest -> init docker with postgrest
- make create_db -> create schema ms-"name-micro-service"
- make drop_db -> delete schema
- migrate_up -> execute migration of create entities
- migrate_down -> execute script of drop table

# Run migrations manual
Up migrations
```
migrate -path db/migration -database "postgresql://postgres:81943301Te@localhost:5009/ms-remision?sslmode=disable" -verbose up
```
Down migrations
```
migrate -path db/migration -database "postgresql://postgres:81943301Te@localhost:5009/ms-remision?sslmode=disable" -verbose down
```


