FROM debian:stable-slim

RUN apt-get update
RUN apt-get install -y tar wget

RUN rm -r /var/cache/apt

ENV NVM_DIR /nvm
ENV NODE_VERSION 12
ENV ZOLA_VERSION 0.9.0

RUN mkdir -p $NVM_DIR
RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.1/install.sh | bash
RUN wget -qO- https://github.com/getzola/zola/releases/download/v${ZOLA_VERSION}/zola-v${ZOLA_VERSION}-x86_64-unknown-linux-gnu.tar.gz |\
	tar xz -C /usr/local/bin

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
