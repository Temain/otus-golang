FROM golang:latest as builder
LABEL maintainer="Otus Aseev A"
WORKDIR /app
COPY . .
RUN GOOS=linux go build -o main ./cmd/calendar_sender

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/configs/config_sender.json ./configs/config_sender.json
COPY --from=builder /app/main .
RUN mkdir logs
EXPOSE 8888
CMD ["./main", "--config=configs/config_sender.json"]