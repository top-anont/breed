include config/env/.env
DB_URL=postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

run: lint
	go run .

lint: download
	golangci-lint run
	
download:
	go mod download


migrate-db-gen:
	migrate create -ext sql -dir ./migrations -seq ${name}

migrate-db:
	migrate -path ./migrations -database "${DB_URL}" up