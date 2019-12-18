up:
	migrate -source file://db/migrations -database "postgres://postgres:example@localhost:5432/golang?sslmode=disable" up

down:
	migrate -source file://db/migrations -database "postgres://postgres:example@localhost:5432/golang?sslmode=disable" down