# FROM golang:1.9-alpine as builder
# WORKDIR /go/src/mphclub-rest-server
# COPY . .
# RUN pwd
#
# FROM scratch
# # Copy our static executable
# COPY --from=builder /go/src/mphclub-rest-server/bin/mphclub-rest-server ./mphclub-rest-server
# COPY --from=builder /go/src/mphclub-rest-server/swagger ./swagger
# CMD ["./mphclub-rest-server"]
#
FROM golang:1.9-alpine as builder
WORKDIR /go/src/mphclub-rest-server
COPY . .
RUN apk update && apk add git && apk add ca-certificates
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./mphclub-rest-server
RUN ls
#CMD ["./mphclub-rest-server"]

FROM scratch
COPY --from=builder /go/src/mphclub-rest-server/mphclub-rest-server ./mphclub-rest-server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/src/mphclub-rest-server/swagger ./swagger
CMD ["./mphclub-rest-server"]
