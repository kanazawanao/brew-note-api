FROM golang:1.20.5-alpine as build-step

WORKDIR /go/api

COPY go.mod go.sum ./
COPY openapi/ ./openapi
COPY src/ ./src

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app/main src/main.go

FROM scratch
COPY --from=build-step /go/api/app/main /go/api/app/main
ENTRYPOINT ["/go/api/app/main"]