export CURRENT_HEAD = $$(git rev-parse HEAD)

run-server:
	@go run api.go

run-client:
	@cd client && go run client.go

new-binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./bin/mphclub-server

#Use this after code push only
docker-build:
	#only have to login (the below command) once per 12 hours
	#@eval `aws ecr get-login --region us-east-1 --no-include-email`
	@docker build -t mphclub_api:${CURRENT_HEAD} -f ./docker/mphclub-rest-server/Dockerfile .

docker-tag:
	@docker tag mphclub_api:${CURRENT_HEAD} mphclub_api:latest 
	@docker tag mphclub_api:latest 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:latest
	@docker tag mphclub_api:latest 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:${CURRENT_HEAD}

docker-push:
	@docker push 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:latest
	@docker push 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:${CURRENT_HEAD}

docker-clean:
	@docker rmi mphclub_api:${CURRENT_HEAD} \
	               mphclub_api:latest \
				   077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:${CURRENT_HEAD} \
				   077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:latest

docker-deploy:
	#only have to apply if the configs change
	#@kubectl apply -f k8s
	@kubectl set image deployments/server-deployment mphclub-api=077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:${CURRENT_HEAD}

export-current:
	@echo ${CURRENT_HEAD}

swagger-html:
	@cd ./swagger && \
	redoc-cli bundle mph-swagger.yaml
