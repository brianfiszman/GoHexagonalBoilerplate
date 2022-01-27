FROM  golang:1.17.1-stretch

WORKDIR /app

COPY . .

RUN `go mod tidy \
  && go build cmd/main.go`

CMD ["./main"]
