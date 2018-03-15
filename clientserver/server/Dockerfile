FROM golang:1.9

WORKDIR /go

COPY . .

RUN go get -d -v ./...

RUN go build main.go

ENTRYPOINT ["/go/main"]

EXPOSE 5300