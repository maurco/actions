FROM docker:19

RUN apk update
RUN apk add --update bash git jq python3
RUN rm /var/cache/apk/*

RUN pip3 install --upgrade pip
RUN pip3 install --upgrade awscli

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
