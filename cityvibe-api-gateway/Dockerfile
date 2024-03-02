FROM golang:1.22-alpine3.19

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -0 ./cmd/ ./cmd/

ENTRYPOINT [ "./cmd/main" ]

