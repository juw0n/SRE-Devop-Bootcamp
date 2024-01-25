
run_postgres:
	# Command to start PostgreSQL on linus ubuntu
	sudo systemctl start postgresql

create_user:
	# Creating a user. If the user already exists, this will throw an error
	@psql -U postgres -c "CREATE USER schooluser WITH PASSWORD 'school123';" || true

create_db:
	# Creating a database. If the database already exists, this will throw an error
	@psql -U postgres -c "CREATE DATABASE studentdb;" || true

	# Grant all privileges of the database to the user
	@psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE studentdb TO schooluser;" || true

drop_db:
	# To drop/delete a database

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


.PHONY: run_postgres create_user create_db drop_db migrate_up migrate_down migrate_fix sqlc test