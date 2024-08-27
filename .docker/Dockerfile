ARG TARGET_OS
ARG TARGET_ARCH

# Build stage

FROM --platform=${TARGET_OS}/${TARGET_ARCH} golang:1.23-alpine3.19 AS build

RUN apk -U add ca-certificates
RUN apk update && apk upgrade && apk add pkgconf git bash build-base sudo
RUN git clone https://github.com/edenhill/librdkafka.git && cd librdkafka && ./configure --prefix /usr && make && make install

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 \
    GO111MODULE=on \
    GOOS=${TARGET_OS} \
    GOARCH=${TARGET_ARCH} \
    go build -mod=vendor -tags musl -ldflags "-extldflags -static -w -s" \
    -o /usr/local/bin/kafka-app ./cmd/main.go

# Final stage

FROM alpine:3.19 

RUN apk add --no-cache libssl3 libcrypto3 ca-certificates tzdata \
    && update-ca-certificates \
    && addgroup -g 1000 "server" \
    && adduser -D -u 1000 -G "server" "server"

COPY schema/ /usr/local/lib/schema/

COPY --from=build /usr/local/bin/kafka-app /usr/local/bin/kafka-app

USER server
