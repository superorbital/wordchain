ARG ARCH=
FROM ${ARCH}golang:1.16-alpine3.13 AS build

RUN apk --no-cache add \
    bash \
    gcc \
    musl-dev \
    openssl

ENV CGO_ENABLED=0

COPY . /build
WORKDIR /build

RUN go get github.com/markbates/pkger/cmd/pkger && \
    pkger -include /data/words.json && \
    go build .

FROM ${ARCH}alpine:3.13 AS deploy

WORKDIR /
COPY --from=build /build/wordchain /

USER 500
EXPOSE 8080

ENTRYPOINT ["/wordchain"]
CMD ["listen"]

