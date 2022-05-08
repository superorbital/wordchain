ARG ARCH=
FROM ${ARCH}golang:1.18-alpine3.15 AS build

RUN apk --no-cache add \
    bash \
    gcc \
    musl-dev \
    openssl

ENV CGO_ENABLED=0

COPY . /build
WORKDIR /build

RUN go install github.com/markbates/pkger/cmd/pkger@latest && \
    pkger -include /data/words.json && \
    go build .

FROM ${ARCH}alpine:3.15 AS deploy

WORKDIR /
COPY --from=build /build/wordchain /

USER 500
EXPOSE 8080

ENTRYPOINT ["/wordchain"]
CMD ["listen"]

