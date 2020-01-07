FROM golang:1.11.0


WORKDIR /go/src/go-grpc-client
COPY . .

# Enable Go Modules.
ENV GO111MODULE=on

RUN go mod tidy
RUN go build ./...
RUN go install -v ./...

# RUN adduser --disabled-password --gecos '' api
# USER api

CMD ["go", "run", "cmd/main.go"]