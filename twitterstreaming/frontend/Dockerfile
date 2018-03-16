FROM golang:1.9
EXPOSE 3000

MAINTAINER brian.coveney

WORKDIR /go
COPY . .

RUN go get -d -v ./...

RUN go build frontend.go

ENTRYPOINT ["/go/frontend"]
