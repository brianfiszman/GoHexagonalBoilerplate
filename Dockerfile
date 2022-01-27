FROM  golang:1.17.1-stretch

WORKDIR /app

COPY . .

RUN ["go", "mod", "tidy"]

CMD ["go","run","cmd/main.go"]
