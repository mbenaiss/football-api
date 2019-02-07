
.PHONY: build clean deploy dev

create:
	gcloud projects create graphql-matches-api --set-as-default
	gcloud services enable cloudfunctions.googleapis.com
	gcloud services enable gcloud.functions.deploy
	export GO111MODULE=on

dependecy: clean
	env GO111MODULE=on go mod init && env GO111MODULE=on go mod vendor

dev:
	go run cmd/main.go

clean:
	rm -rf ./vendor ./go.mod ./go.sum

deploy: clean dependecy
	gcloud functions deploy matches --runtime go111 --entry-point Matches --trigger-http --project matches  
