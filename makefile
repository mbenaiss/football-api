.PHONY: dev

dev:
	go run cmd/main.go

deploy-gcp:
	cd gcp && make deploy