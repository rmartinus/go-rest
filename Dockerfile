FROM golang:1.13 as builder
COPY . /app
WORKDIR /app/

RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -installsuffix cgo -o service cmd/service/main.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/service /app/api/*.yaml /app/
CMD ["./service"]
