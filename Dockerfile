FROM ubuntu:latest

MAINTAINER takecy

RUN apt-get -y update \
 && apt-get install -y wget

WORKDIR /usr/local/webp
RUN wget https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-0.5.0-linux-x86-64.tar.gz \
 && tar xfvz libwebp-0.5.0-linux-x86-64.tar.gz

ENV PATH $PATH:/usr/local/webp/libwebp-0.5.0-linux-x86-64/bin
RUN cwebp -h \
 && dwebp -h
