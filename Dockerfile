FROM golang:1-alpine as build

ARG action

WORKDIR /app
COPY $action/src src
COPY toolkit toolkit
COPY go.mod go.sum ./

RUN go build -v -o entrypoint ./src

FROM alpine

COPY --from=build /app/entrypoint /usr/local/bin/
ENTRYPOINT ["entrypoint"]
