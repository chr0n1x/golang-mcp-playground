FROM golang:1.24.2

RUN apt update && apt install -y entr

ADD . /go/src/github.com/chr0n1x/golang-mcp-playground
WORKDIR /go/src/github.com/chr0n1x/golang-mcp-playground


RUN go mod tidy

CMD ["go", "run", "main.go"]
