FROM golang:1-alpine as build

ARG action
WORKDIR /app
COPY $action/src $action/go.mod $action/go.sum ./

RUN go build -v -o entry ./...

FROM alpine

RUN apk update
RUN apk add --no-cache --upgrade \
	bash \
	git \
	libstdc++ \
	libc6-compat \
	tar \
	wget

COPY --from=build /app/entry /usr/local/bin/

ENTRYPOINT ["entry"]
