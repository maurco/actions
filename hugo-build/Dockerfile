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

RUN wget -qO /usr/local/bin/hugo-build https://github.com/maurerlabs/github-actions/raw/master/hugo-build/bin/hugo-build
RUN chmod +x /usr/local/bin/hugo-build

# only if local binary exists (for testing)
COPY README.md bin/* /usr/local/bin

ENTRYPOINT ["hugo-build"]
