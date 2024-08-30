###----- Builder -----###
FROM golang:1.21 AS builder

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

###----- Production -----###
FROM debian:bookworm-slim AS production

WORKDIR /app

ENV GIN_MODE=release

COPY --from=builder /app/main /app/main

EXPOSE ${PORT}

CMD ["/app/main"]

