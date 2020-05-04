FROM golang:latest as builder
LABEL maintainer="Otus Aseev A"
WORKDIR /app
COPY .. .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/rotation_api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/configs/config.json ./configs/config.json
COPY --from=builder /app/main .
RUN mkdir logs
CMD ["./main", "--config=configs/config.json"]

