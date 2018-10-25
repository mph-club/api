gen_proto:
	@cd api-generated && \
	protoc -I/usr/local/include -I. \
	  -I$(GOPATH)/src \
	  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	  --go_out=plugins=grpc:. \
	  mphclubapispec.proto && \
	protoc -I/usr/local/include -I. \
	  -I$(GOPATH)/src \
	  -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	  --grpc-gateway_out=logtostderr=true:. \
	  mphclubapispec.proto

		@echo "generated"

gen_server_tls:
	@openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt -subj "/CN=mphclub.biz/O=mphclub"

gen_client_tls:
	@openssl genrsa -out client-test/client.key 2048
	@openssl req -new -x509 -sha256 -key client-test/client.key -out client-test/client.crt -days 3650

run/server:
	@go run api.go

run/client:
	@go run client-test/client.go

new-binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o docker/go-app-prod/mphclub-server

push-latest:
		@docker push 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:latest

docker-build-api:
	@docker build -t mphclub_api -f ./Dockerfile .
	@docker tag mphclub_api:latest 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:latest
	@docker push 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_api:latest

docker-build-pg:
	@docker build -t mphclub_pg -f ./docker/postgres/Dockerfile ./docker/postgres
	@docker tag mphclub_pg:latest 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_pg:latest
	@docker push 077003688714.dkr.ecr.us-east-1.amazonaws.com/mphclub_pg:latest

docker-script:
	@sh docker-prodbuild.sh

proto-from-oapi:
	@openapi2proto -spec mph-swagger.yaml -annotate -out api-generated/mphclubapispec.proto && \
	redoc-cli bundle mph-swagger.yaml && \
	mv redoc-static.html swagger/index.html

test-json-userdata:
	@echo "makes authenticated call \n"
	@curl -X POST -kv https://gateway.mphclub.biz/v1/userdata -H "Content-Type: text/plain" -H "Authorization: eyJraWQiOiJtS1wvSjRWV1UzcTI1WkNrVjN1TGZSQXZIQnJ4OHR5YWNEbVU5c2xNaEUrOD0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIwNmRmNzI0OS1mZWQyLTQ3ZGItOTA4My01YTU4NDY3NzU1ZTMiLCJldmVudF9pZCI6IjA3MTcyMWQwLWMwMTYtMTFlOC1hNTg3LTFkMWYwZDM5NmIwYSIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE1Mzc4MDYxNDcsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC51cy1lYXN0LTEuYW1hem9uYXdzLmNvbVwvdXMtZWFzdC0xX1JLc0pBOXd1YSIsImV4cCI6MTUzNzgwOTc0NywiaWF0IjoxNTM3ODA2MTQ3LCJqdGkiOiJmODQwMTk4YS0zOWVjLTQyZTYtYTE5Yi0wNmI4MTJmZTVmNDQiLCJjbGllbnRfaWQiOiIyODk1c3BtazlvNnZ1cWN0YmVkYzB2OWYzOCIsInVzZXJuYW1lIjoiMDZkZjcyNDktZmVkMi00N2RiLTkwODMtNWE1ODQ2Nzc1NWUzIn0.KRyUdXI4ilkmYMubb3aKJS0v5fYpSkla4osmsjrcg-7a51VEl5B13ilqNNU-gRegj3h5xIvzqRHCdhn-1OVZXg2iKXj9mpBlBoJX2Dg4MinLGvDrKi3Y8loWkktBBGUQRIyZ3ne970TCFSUpKdHAZ2J5A5T0-6_La9LwsM711S2Vc8khNfwcu1Zlro3SYS4vpCUNQ3qLj6tw9F7Fx5VxYUV-kQ14l3SUBqBtCUQ5pkLYXmHrMOPl9ztXe4QHHjVt52Wj4OB-VR5wb9OhCb0jnPbSrOX7I8bmQQl4iyGmksry_kqvi2UPsXjfAcyraUI2SUZK2nUudq11rX19_PDQig" -d '{ "userData": { "userAccount": {"sub" : "hello"} } }'

test-json-vehiclelist:
	@echo "gets 5 sedans \n"
	@curl -X GET -kv https://gateway.mphclub.biz/v1/vehicle/getVehicleList/sedan/5
