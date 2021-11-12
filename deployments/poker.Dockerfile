FROM golang:1.17.3-alpine3.14 as builder

WORKDIR poker
ENV GO111MODULE=on CGO_ENABLED=1
RUN apk add --no-cache git openssh-client

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -ldflags '-w -s' -o ./bin/app ./cmd/poker

FROM alpine:3.7

COPY . .

COPY --from=builder /go/poker/bin/app /app

RUN chmod +x /app

CMD ["/app"]
