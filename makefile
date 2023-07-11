POSTGRESQL_URL=postgres://postgres:secret@localhost:5432/go-todo?sslmode=disable
CONTAINER_NAME=iamtbay-todo-db-1
run: build
	@ ./todo
build:
	@ go build -o todo ./cmd/*.go
dbmigrateup:
	@ migrate -database ${POSTGRESQL_URL} -path db/migrations up
dbmigratedown:
	@ migrate -database ${POSTGRESQL_URL} -path db/migrations down

dockerup:
	docker-compose -f docker-compose.yaml up
dockerstop:
	docker stop ${CONTAINER_NAME}
dockerstart:
	docker start ${CONTAINER_NAME}
restart: dockerstop dockerstart run
