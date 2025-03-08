FROM golang:1.23.3-alpine AS builder
WORKDIR /miner

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o miner main.go

FROM  alpine:latest
WORKDIR /root/miner

COPY /miner/redis-cluster ./redis-cluster
COPY --from=builder /miner/miner ./
COPY --from=builder /miner/config.yml ./

EXPOSE 9090
CMD ["./miner"]
