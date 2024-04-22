sync-deps:
	go mod tidy
	go mod vendor

consumer:
	go run main.go consumer -e=local