FROM golang:1.19 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o ./build/rpc_server/main ./cmd/rpc_server/main.go

FROM scratch

WORKDIR /

COPY --from=builder /app/build/rpc_server/main /rpc_server
COPY --from=builder /app/utl/config/config.local.yaml /utl/config/config.local.yaml

EXPOSE 8080:8080

ENTRYPOINT ["/rpc_server"]
