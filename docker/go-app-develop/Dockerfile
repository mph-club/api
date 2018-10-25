FROM golang:1.9-alpine
WORKDIR /go/src/mphclub-rest-server
COPY . .
RUN apk add --no-cache curl
RUN curl -L -s https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -o $GOPATH/bin/dep
RUN chmod +x $GOPATH/bin/dep
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o mphclub-rest-server
CMD ["mphclub-rest-server"]
