FROM golang:1.24-alpine
EXPOSE 8080
WORKDIR /usr/src/app

COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go test ./...
RUN go build -o /usr/local/bin/app cmd/main.go

CMD ["/usr/local/bin/app"]
