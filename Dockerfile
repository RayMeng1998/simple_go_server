FROM golang

WORKDIR /mzy/go_http_server

EXPOSE 8080

CMD ["go", "run", "go_server.go"]

COPY . .
