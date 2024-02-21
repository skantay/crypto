migrateup:
	migrate -path=./migrations -database="postgres://user:pass@localhost:5432?sslmode=disable" up

migratedown:
	migrate -path=./migrations -database="postgres://user:pass@localhost:5432?sslmode=disable" down

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