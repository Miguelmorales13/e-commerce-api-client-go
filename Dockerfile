FROM golang:latest as builder
RUN apt-get update
ENV GOOS=linux \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOARCH=amd64 \
    PORT=:8000 \
    DB_HOST=host \
    DB_USER=postgres \
    DB_PASSWORD=root \
    DB_NAME=test_go \
    DB_PORT=5432 \
    DB_SSMODE=disable
WORKDIR /go/bin/app
COPY go.mod .
RUN go mod download
#RUN mkdir .env
COPY . .
RUN go build crud.go

FROM scratch
COPY --from=builder /go/bin/app .
ENTRYPOINT ["./crud"]
