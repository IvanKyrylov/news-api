FROM golang:1.18-alpine3.16 as builder
WORKDIR /usr/local/go/src/
COPY . .
RUN go clean --modcache
RUN go build -mod=readonly -o app cmd/news-api/main.go

FROM alpine:3.16
COPY --from=builder /usr/local/go/src/app /
COPY --from=builder /usr/local/go/src/config.yml /

CMD ["/app"]