FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/MarkCherepovskyi/KeyStorage
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/KeyStorage /go/src/gitlab.com/MarkCherepovskyi/KeyStorage


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/KeyStorage /usr/local/bin/KeyStorage
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["KeyStorage"]
