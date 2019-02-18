FOOTBALL_API_KEY=[[YOUR_API_KEY]]

.PHONY: dependecy clean deploy dev
dependecy: clean
	env GO111MODULE=on go mod init  && env GO111MODULE=on go mod vendor

clean:
	rm -rf ./vendor ./go.mod ./go.sum

deploy: clean dependecy
	gcloud functions deploy fn --runtime go111 \
	 --entry-point Matches --trigger-http
	 --set-env-vars FOOTBALL_API_KEY=${FOOTBALL_API_KEY}