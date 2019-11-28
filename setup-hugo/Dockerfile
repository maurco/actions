FROM golang:1-alpine as build

WORKDIR /app
COPY src src
COPY go.mod go.sum ./

RUN go build -v -o setup-hugo ./...

FROM alpine

COPY --from=build /app/setup-hugo /usr/local/bin/
ENTRYPOINT ["setup-hugo"]
