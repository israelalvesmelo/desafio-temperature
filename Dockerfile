FROM golang:1.23-alpine as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/api

FROM scratch
COPY --from=builder /app/env.json .
COPY --from=builder /app/server .

CMD ["./server"]