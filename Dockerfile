FROM golang:latest

RUN go version
ENV GOPATH=/

COPY go.mod .
COPY go.sum .

COPY ./ ./

RUN go mod download
RUN go build -o main ./cmd/main.go
CMD ["./main"]