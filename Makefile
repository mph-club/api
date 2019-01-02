IMAGE_NAME = mphclub_api
CURRENT_HEAD = $$(git rev-parse HEAD)
ACCOUNT_ID = $$(aws sts get-caller-identity --output text --query 'Account')
REGISTRY_URL = ${ACCOUNT_ID}.dkr.ecr.us-east-1.amazonaws.com
KUBE_DEPLOYMENT = deployments/server-deployment
KUBE_CONTAINER_NAME = mphclub-api
STAGE = dev

run-server:
	@go run api.go

run-client:
	@cd client && go run client.go

new-binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./bin/mphclub-server

#Use this after code push only
docker-build:
	@docker build --force-rm \
	              --tag ${IMAGE_NAME}:${CURRENT_HEAD} \
	              --file ./docker/mphclub-rest-server/Dockerfile .

docker-tag:
	@docker tag ${IMAGE_NAME}:${CURRENT_HEAD} ${IMAGE_NAME}:latest
	@docker tag ${IMAGE_NAME}:latest ${REGISTRY_URL}/${IMAGE_NAME}:latest
	@docker tag ${IMAGE_NAME}:latest ${REGISTRY_URL}/${IMAGE_NAME}:${CURRENT_HEAD}
	@docker tag ${IMAGE_NAME}:${CURRENT_HEAD} ${IMAGE_NAME}:${STAGE}
	@docker tag ${IMAGE_NAME}:${STAGE} ${REGISTRY_URL}/${IMAGE_NAME}:${STAGE}
	@docker tag ${IMAGE_NAME}:${STAGE} ${REGISTRY_URL}/${IMAGE_NAME}:${CURRENT_HEAD}

docker-push:
	# only have to login (the below command) once per 12 hours
	# @eval `aws ecr get-login --no-include-email`
	@docker push ${REGISTRY_URL}/${IMAGE_NAME}:latest
	@docker push ${REGISTRY_URL}/${IMAGE_NAME}:${CURRENT_HEAD}
	@docker push ${REGISTRY_URL}/${IMAGE_NAME}:${STAGE}
	# @docker logout ${REGISTRY_URL}

docker-clean:
	@docker rmi --force \
	            ${IMAGE_NAME}:${CURRENT_HEAD} \
	            ${IMAGE_NAME}:latest \
				${REGISTRY_URL}/${IMAGE_NAME}:${CURRENT_HEAD} \
				${REGISTRY_URL}/${IMAGE_NAME}:latest
	# @docker image prune --force

docker-deploy:
	#only have to apply if the configs change
	#@kubectl apply -f k8s
	@kubectl set image ${KUBE_DEPLOYMENT} ${KUBE_CONTAINER_NAME}=${REGISTRY_URL}/${IMAGE_NAME}:${CURRENT_HEAD}

export-current:
	@echo ${CURRENT_HEAD}

account-id:
	@echo ${ACCOUNT_ID}

registry-url:
	@echo ${REGISTRY_URL}

swagger-html:
	@cd ./swagger && \
	redoc-cli bundle mph-swagger.yaml
