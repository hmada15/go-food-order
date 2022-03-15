migrate-up:
	migrate -path database/migrations -database "mysql://root@tcp(localhost:3306)/go-food-order" up
migrate-down:
	migrate -path database/migrations -database "mysql://root@tcp(localhost:3306)/go-food-order" down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...