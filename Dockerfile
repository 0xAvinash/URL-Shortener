FROM golang:1.26.2 AS build

WORKDIR /app

RUN apt-get update && apt-get install -y \
    librdkafka-dev gcc pkg-config

ENV GOTOOLCHAIN=auto

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/main.go


FROM debian:bookworm-slim
WORKDIR /root/

RUN apt-get update && apt-get install -y librdkafka1 && rm -rf /var/lib/apt/lists/*

COPY --from=build /app/main .

RUN chmod +x main

EXPOSE 3000

CMD ["./main"]