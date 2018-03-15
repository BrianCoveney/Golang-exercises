FROM golang:1.9

MAINTAINER brian.coveney

WORKDIR /go
COPY . .

RUN go get -d -v ./...
RUN go build twitter-service.go

ENTRYPOINT ["/go/twitter-service"]
