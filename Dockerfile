FROM golang:1.14-alpine AS build-env

# install build tools
RUN apk add --no-cache \
  git \
  build-base \
  libc6-compat \
  ca-certificates \
  tini

RUN go get -u -v \
  github.com/oxequa/realize && \
  github.com/jinzhu/gorm && \
  github.com/jinzhu/gorm/dialects/postgres

WORKDIR /go/src
COPY . .
RUN go build -o server main.go

EXPOSE 8080
ENTRYPOINT ["/sbin/tini", "--"]