FROM golang:1.9-alpine as builder
WORKDIR /go/src/mphclub-rest-server
RUN apk update && apk add git && apk add ca-certificates && apk add --no-cache curl
RUN apk add tzdata
RUN curl -L -s https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -o $GOPATH/bin/dep
RUN chmod +x $GOPATH/bin/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./mphclub-rest-server

FROM scratch
COPY --from=builder /go/src/mphclub-rest-server/mphclub-rest-server ./mphclub-rest-server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo/America/New_York /usr/share/zoneinfo/America/New_York
COPY --from=builder /go/src/mphclub-rest-server/swagger ./swagger
CMD ["./mphclub-rest-server"]
