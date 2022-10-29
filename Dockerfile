FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/tokend/key_storage
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/key_storage /go/src/gitlab.com/tokend/key_storage


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/key_storage /usr/local/bin/key_storage
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["key_storage"]
