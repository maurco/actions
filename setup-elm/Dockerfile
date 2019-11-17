FROM alpine:3

RUN apk update
RUN apk add --update bash wget

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
