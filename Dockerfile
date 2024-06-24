FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN ls
RUN pwd
RUN go build -o main ./cmd/server/server.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env
EXPOSE 8080
CMD [ "/app/main" ]