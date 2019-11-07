FROM alpine:3.10

RUN apk update
RUN apk upgrade
RUN apk add bash tar wget make
RUN rm /var/cache/apk/*

RUN wget -qO- https://github.com/gohugoio/hugo/releases/download/v0.59.1/hugo_extended_0.59.1_Linux-64bit.tar.gz | tar xzf -
RUN mv hugo /usr/bin

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
