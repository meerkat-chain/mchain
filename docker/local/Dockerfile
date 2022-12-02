FROM golang:1.18-alpine AS builder

WORKDIR /mchain

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o mchain main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /mchain

COPY --from=builder /mchain/mchain ./
COPY ./docker/local/mchain.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478

ENTRYPOINT ["./mchain.sh"]
