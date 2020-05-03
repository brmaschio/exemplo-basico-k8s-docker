FROM golang:1.10-alpine3.8 as builder

WORKDIR /go/src/app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"

FROM scratch

WORKDIR /go/src/app

COPY --from=builder /go/src/app/app .
COPY --from=builder /go/src/app/index.html .

CMD ["./app"]