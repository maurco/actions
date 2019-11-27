FROM maurerlabs/action-hugo-build:latest

ARG HUGO_VERSION=0.59.1
RUN wget -qO- https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_extended_${HUGO_VERSION}_Linux-64bit.tar.gz |\
	tar xz -C /usr/local/bin
