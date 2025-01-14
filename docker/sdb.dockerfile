FROM golang:1.20-alpine AS build
RUN apk add --no-cache --update gcc g++ git
ARG SRCROOT=.
ENV GOPATH /app
RUN mkdir -p /app/output
WORKDIR /app/src

COPY ./ /app/src/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -buildmode=exe -o /app/output/ $SRCROOT/cmd/sdb

FROM alpine:latest
COPY --from=build /app/output/sdb /usr/local/bin

RUN wget -c https://github.com/jupiter/parquet2json/releases/download/v2.0.2/parquet2json-i686-unknown-linux-musl.tar.gz -O - | tar xz -C /usr/local/bin

ENTRYPOINT ["/usr/local/bin/sdb"]
