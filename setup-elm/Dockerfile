FROM golang:1-alpine as build

WORKDIR /app
COPY src src
COPY go.mod go.sum ./

RUN go build -v -o setup-elm ./...

FROM alpine

COPY --from=build /app/setup-elm /usr/local/bin/
ENTRYPOINT ["setup-elm"]
