FROM golang:latest as builder
LABEL maintainer="Otus Aseev A"
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/calendar_api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/configs/config_api.json ./configs/config_api.json
COPY --from=builder /app/main .
RUN mkdir logs
EXPOSE 8888
CMD ["./main", "grpc_server", "--config=configs/config_api.json"]

