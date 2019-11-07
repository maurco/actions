FROM alpine:3.10

RUN apk update
RUN apk upgrade
RUN apk add --update \
	libstdc++ \
	libc6-compat \
	git \
	bash \
	tar \
	wget \
	make

RUN rm /var/cache/apk/*

ENV VERSION 0.59.1
RUN wget -qO- https://github.com/gohugoio/hugo/releases/download/v${VERSION}/hugo_extended_${VERSION}_Linux-64bit.tar.gz | tar xzf -
RUN mv hugo /usr/local/bin

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
