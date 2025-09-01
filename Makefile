redis-docker:
	# * Pull the latest Redis official image.
	docker pull redis:latest
	# * Launch the Redis container.
	docker run -d --name my-redis -p 6379:6379 redis

install:
	go mod tidy

run:
	go run ./src/main.go
