FROM golang:latest as builder
LABEL maintainer="Otus Aseev A"
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/calendar_scheduler

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/configs/config_scheduler.json ./configs/config_scheduler.json
COPY --from=builder /app/main .
RUN mkdir logs
EXPOSE 8888
CMD ["./main", "--config=configs/config_scheduler.json"]

