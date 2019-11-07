FROM alpine:3.10

RUN apk update
RUN apk upgrade
RUN apk add bash python3
RUN pip3 install --upgrade awscli==1.16.*
RUN rm /var/cache/apk/*

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
