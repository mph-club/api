FROM golang:1.9-alpine as builder
WORKDIR /go/src/mphclub-rest-server
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./bin/mphclub-rest-server

FROM scratch
# Copy our static executable
COPY --from=builder /go/src/mphclub-rest-server/bin/mphclub-rest-server ./mphclub-rest-server
COPY --from=builder /go/src/mphclub-rest-server/swagger ./swagger
CMD ["./mphclub-rest-server"]
