POSTGRES_URL="postgres://user:pass@localhost:5432/domain?sslmode=disable"

migrateup:
	migrate -path=./migrations -database=$(POSTGRES_URL) up

migratedown:
	migrate -path=./migrations -database=$(POSTGRES_URL) down

dockershell:
	docker exec -it crypto_postgres_1 bash

dockerup:
	docker-compose up --build

dockerdown:
	docker-compose down

go-test:
	go test ./... -v --cover

git-push:
	git add .
	git commit -m "uploading"
	git push
