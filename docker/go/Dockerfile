###----- Builder -----###
FROM golang:1.22 AS builder

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -C cmd/sampleTodoApp -v -o /app/main

###----- Production -----###
FROM debian:bookworm-slim AS production

WORKDIR /app

ENV GIN_MODE=release

COPY --from=builder /app/main /app/server

EXPOSE ${PORT}

CMD ["/app/server"]

