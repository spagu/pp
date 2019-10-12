  
FROM golang:1.10-alpine as go

RUN apk update \
	&& apk add git \
	&& rm -rf /var/cache/apk/*

COPY . /go/src/pp
WORKDIR /go/src/pp

RUN go get ./...
RUN go build
RUN ./pp 

RUN go test .
RUN go build -ldflags="-w -s -X main.version=$VERSION" .

FROM alpine:latest
COPY --from=go /go/src/pp/pp /bin/pp
RUN mkdir /mount
WORKDIR /mount
ENTRYPOINT ["pp"]
CMD ["--help"]
