FROM golang:1.20.5-alpine
WORKDIR /migrations
COPY . .
RUN apk add --no-cache gcc libc-dev
RUN go install github.com/rubenv/sql-migrate/...@latest
