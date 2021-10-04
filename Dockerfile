FROM  golang:1.17.1-stretch

WORKDIR /app

COPY . .

VOLUME ["/app"]

CMD ["go","run","pkg/main.go"]
