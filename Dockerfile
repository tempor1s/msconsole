FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/msconsole/main.go" --command='./main' -log-prefix=false