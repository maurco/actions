FROM golang:1-alpine as build

ARG action
WORKDIR /app
COPY $action/src src
COPY $action/go.mod $action/go.sum ./

RUN go build -v -o entry ./...

FROM alpine

COPY --from=build /app/entry /usr/local/bin/
ENTRYPOINT ["entry"]
