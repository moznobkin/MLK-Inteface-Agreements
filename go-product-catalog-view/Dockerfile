FROM golang:1.12.7-alpine3.10 AS builder
RUN apk add git

RUN mkdir -p /go/src/github.com/moznobkin/product-catalog-view
COPY . /go/src/github.com/moznobkin/product-catalog-view
WORKDIR /go/src/github.com/moznobkin/product-catalog-view
ENV GO111MODULE=on
RUN go build -v -o /go/bin/product-catalog-view

FROM alpine:3.10
LABEL maintainer="Michael Oznobkin <oznobkin@gmail.com>"
COPY --from=builder /go/bin/product-catalog-view /usr/bin/product-catalog-view
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/product-catalog-view"]


