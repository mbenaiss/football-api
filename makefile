PROJETC_ID=football-231615
USER_EMAIL=mahmoud.b.aissa@gmail.com

.PHONY: build clean deploy dev

create:
	gcloud projects create ${PROJETC_ID} --set-as-default

permission:
	gcloud services enable cloudfunctions.googleapis.com --project=${PROJETC_ID}
	gcloud projects add-iam-policy-binding ${PROJETC_ID} \
  --member='user:${USER_EMAIL}' \
  --role 'roles/owner'
	export GO111MODULE=on

dependecy: clean
	env GO111MODULE=on go mod init  && env GO111MODULE=on go mod tidy && env GO111MODULE=on go mod vendor

dev:
	go run cmd/main.go

clean:
	rm -rf ./vendor ./go.mod ./go.sum

deploy: clean dependecy
	gcloud functions deploy function --runtime go111 \
	 --entry-point Matches --trigger-http
