FROM golang:1.12.7-alpine3.10 AS builder
RUN mkdir -p /go/src/github.com/moznobkin/go-product-catalog-view
COPY ./go-product-catalog-view /go/src/github.com/moznobkin/go-product-catalog-view
WORKDIR /go/src/github.com/moznobkin/go-product-catalog-view
RUN go build -v -o /go/bin/go-product-catalog-view

FROM alpine:3.10
LABEL maintainer="Michael Oznobkin <oznobkin@gmail.com>"
COPY --from=builder /go/bin/go-product-catalog-view /usr/bin/go-product-catalog-view
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/go-product-catalog-view"]