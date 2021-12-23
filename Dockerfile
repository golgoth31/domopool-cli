# Get certificates for https conexions
FROM alpine:latest as certs
RUN apk add -U --no-cache ca-certificates

# build UI
FROM node:12-alpine as react-build
WORKDIR /usr/src/app
COPY ./web .
RUN yarn install && yarn build

# build binary
FROM golang:1.16 as golang-build
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org
WORKDIR /go/src/github.com/golgoth31/domopool
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=react-build /usr/src/app/build ./web/build
RUN make build_local

FROM scratch
VOLUME /data /config
ENV PATH=/bin
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=golang-build /go/src/github.com/golgoth31/domopool/domopool /domopool
EXPOSE 8080
ENTRYPOINT ["/domopool"]
