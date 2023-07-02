FROM golang:1.20.5-alpine as build-step
WORKDIR /go/tripig
COPY . /go/tripig
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app/main src/main.go

FROM scratch
COPY --from=build-step /go/tripig/app/main /go/tripig/app/main
ENTRYPOINT ["/go/tripig/app/main"]