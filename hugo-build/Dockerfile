FROM golang:1-alpine AS build

WORKDIR /app
COPY . .
RUN go build -v -o /usr/local/bin/hugo-build ./...

FROM alpine

RUN apk update
RUN apk add --update \
	git \
	libstdc++ \
	libc6-compat \
	tar \
	wget

ARG HUGO_VERSION=0.59.1
RUN wget -qO- https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_extended_${HUGO_VERSION}_Linux-64bit.tar.gz |\
	tar xz -C /usr/local/bin

COPY --from=build /usr/local/bin/hugo-build /usr/local/bin
# COPY bin/hugo-build /usr/local/bin
ENTRYPOINT ["hugo-build"]
