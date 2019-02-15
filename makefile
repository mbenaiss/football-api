.PHONY: dependecy clean deploy dev

dependecy: clean
	env GO111MODULE=on go mod init  && env GO111MODULE=on go mod tidy

dev:
	go run cmd/main.go

clean:
	rm -rf ./vendor ./go.mod ./go.sum
