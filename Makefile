## pull the postgres image
pull_postgres:
	docker pull postgres:16.2-alpine3.18

postgres:
	# Command to start PostgreSQL on linus ubuntu
	docker run --name postgres16 -e POSTGRES_USER=schooluser -e POSTGRES_PASSWORD=school123 -p 5432:5432 -d postgres:16.2-alpine3.18

create_db:
	# Creating a database
	docker exec -it postgres16 createdb --username=schooluser  --owner=schooluser studentdb

drop_db:
	# To drop/delete a database
	docker exec -it postgres16 dropdb -U schooluser studentdb

migrate_up:
	# To execute the SQL statements in the migration files and create the define tables
	migrate -path database/migration/ -database "postgresql://schooluser:school123@localhost:5432/studentdb?sslmode=disable" -verbose up

migrate_down:
	# To rollback or revert back a migration to the formal state
	migrate -path database/migration/ -database "postgresql://schooluser:school123@localhost:5432/studentdb?sslmode=disable" -verbose down

migrate_fix:
	# To Resolve a Migration "Dirty database version 1. Fix and force version" Errors 
	migrate -path database/migration/ -database "postgresql://schooluser:school123@localhost:5432/studentdb?sslmode=disable" force <version number of the migrate>

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -source=/home/juwon/Desktop/cloudComputingLessons/SRE-Devop-Bootcamp/database/sqlc/querier.go -destination=mocks/student_mock.go -package=mocks

build:
	docker build -t student-go-api:1.0.0 .

run_api:
	# update the host IP to the ip of the postgresql
	docker run --name student-api -p 8080:8000 -e DB_SOURCE="postgresql://schooluser:school123@localhost:5432/studentdb?sslmode=disable" student-go-api:v1

.PHONY: pull_postgres postgres create_user create_db drop_db migrate_up migrate_down migrate_fix sqlc test server mock build run_api