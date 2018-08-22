# STEP 1 build executable binary
FROM golang:alpine as builder

# Install SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates

# Create appuser
RUN adduser -D -g '' appuser
WORKDIR /go/bin/mphclub-golang-api
COPY . .
#RUN chmod +x /go/bin/mphclub-golang-api

# STEP 2 build a small image
# start from scratch
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable
COPY --from=builder /go/bin/mphclub-golang-api ./
#RUN chmod +x /app/mphclub-golang-api
CMD ["/mphclub-golang-api"]
#
