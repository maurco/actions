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

COPY bin/hugo-build-alpine /usr/local/bin/hugo-build
ENTRYPOINT ["hugo-build"]
