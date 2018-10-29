run/server:
	@go run api.go

new-binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o docker/go-app-prod/mphclub-server

#Use this after code push only
docker-build-api:
	@export CURRENT_HEAD=$(git rev-parse HEAD)
	@echo $${CURRENT_HEAD}
	#only have to login (the below command) once per 12 hours
	@eval `aws ecr get-login --region us-east-1 --no-include-email`
	@docker build -t mphclub_api -f ./docker/go-app-develop/Dockerfile .
	@docker tag mphclub_api:latest 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:latest
	@docker tag mphclub_api:latest 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:$${CURRENT_HEAD}
	@docker push 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:latest
	@docker push 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:$${CURRENT_HEAD}
	@kubectl set image deployments/server-deployment mphclub-api=077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:$${CURRENT_HEAD}

swagger-html:
	@cd ./swagger && \
	redoc-cli bundle mph-swagger.yaml
