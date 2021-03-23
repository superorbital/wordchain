FROM golang:1.16-alpine3.13 AS build

RUN apk --no-cache add \
    bash \
    gcc \
    musl-dev \
    openssl
RUN mkdir -p /go/src/github.com/superorbital/wordchain
WORKDIR /go/src/github.com/superorbital/wordchain
ADD . /go/src/github.com/superorbital/wordchain
RUN go get github.com/markbates/pkger/cmd/pkger && \
    pkger -include /data/words.json && \
    go build -mod=vendor --ldflags '-linkmode external -extldflags "-static"' .

FROM alpine:3.13 AS deploy

WORKDIR /
COPY --from=build /go/src/github.com/superorbital/wordchain/wordchain /

USER 500

ENTRYPOINT ["/wordchain"]
CMD ["random"]
