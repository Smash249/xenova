
FROM golang:1.25.6-alpine AS builder

WORKDIR /app


COPY go.mod ./


COPY . .

RUN go mod tidy && go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest

WORKDIR /app


COPY --from=builder /app/main .


COPY --from=builder /app/config ./config




EXPOSE 8099


CMD ["./main"]
