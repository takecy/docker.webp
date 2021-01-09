FROM alpine:3.12

LABEL maintainer=takecy
WORKDIR /usr/local/webp

ENV WEBP_VERSION 1.1.0-r0

RUN apk update\
 && apk add --no-cache ca-certificates tzdata openssl libwebp-tools=${WEBP_VERSION} bash

RUN cwebp -h \
 && dwebp -h

WORKDIR /usr/local/img
