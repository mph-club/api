FROM golang:1.9-alpine
COPY root /
RUN apk add --no-cache ca-certificates \
        dpkg \
        gcc \
        git \
        musl-dev \
        bash
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && chmod +x /entrypoint.sh
RUN go get -u github.com/tockins/realize
RUN go get -u github.com/kataras/iris
WORKDIR $GOPATH
ENTRYPOINT ["/entrypoint.sh"]
CMD realize start
