setup:
	docker-compose -f tools/development/docker-compose.yaml up -d

sync-deps:
	go mod tidy
	go mod vendor

consumer:
	go run main.go consumer -e=local

consumer-dev:
	go run main.go consumer -e=dev

consumer-prd:
	go run main.go consumer -e=prod

consumer-sandbox:
	go run main.go consumer -e=sandbox