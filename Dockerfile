FROM golang:1.14



ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on
EXPOSE 8080


CMD ["go", "run","main.go"]