FROM golang:alpine

WORKDIR /app

COPY ./ ./
EXPOSE 8000

CMD ["/usr/local/go/bin/go", "./main.go"]
