FROM golang:1.23-alpine AS builder


WORKDIR /app
COPY go.mod ./

COPY . .

RUN go mod tidy && go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8082

RUN addgroup -g 1001 -S appuser && \
    adduser -S appuser -u 1001 -G appuser

RUN chown appuser:appuser /root/main
USER appuser

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8082/health || exit 1

CMD ["./main"]
