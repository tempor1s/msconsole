FROM golang:alpine

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 5000

ENTRYPOINT CompileDaemon --build="go build ./server/main.go" --command=./main -log-prefix=false