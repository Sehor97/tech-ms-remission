postgres:
	docker run --name ms-remission -p 5009:5432 -e POSTGRES_PASSWORD=81943301Te -d postgres

create_db:
	docker exec -it ms-remission createdb --username=postgres --owner=postgres ms-remission

drop_db:
	docker exec -it ms-remission dropdb ms-remission

migrate_up:
	migrate -path db/migration -database "postgresql://postgres:81943301Te@localhost:5009/ms-remission?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://postgres:81943301Te@localhost:5009/ms-remission?sslmode=disable" -verbose down


.PHONY: postgres create_db drop_db migrate_up migrate_down
