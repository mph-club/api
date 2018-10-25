run/server:
	@go run api.go

new-binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o docker/go-app-prod/mphclub-server

push-latest:
		@docker push 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:latest

docker-build-api:
	@docker build -t mphclub_api -f ./docker/go-app-develop/Dockerfile .
	@docker tag mphclub_api:latest 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:latest
	@docker push 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:latest
	@sh push_and_make_token.sh

swagger-html:
	@cd ./swagger && \
	redoc-cli bundle mph-swagger.yaml
