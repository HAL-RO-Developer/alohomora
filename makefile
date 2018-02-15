create:
	cp dbconfig.yml.template dbconfig.yml
	
migrate:
	go run cmd/migrate.go

run:
	env lavender=production go run main.go

docker/build:
	env GOOS=linux env GOARCH=amd64 env CGO_ENABLED=0 go build -o ./cmd/main main.go
	docker build -t makki0205/alohomora .

docker/run:
	docker run -it --rm makki0205/alohomora

docker/push:
	docker push makki0205/alohomora
