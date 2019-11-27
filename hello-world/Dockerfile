FROM golang:1-alpine as build

WORKDIR /app
COPY src src
COPY go.mod go.sum ./

RUN go build -v -o hello-world ./...

FROM alpine

COPY --from=build /app/hello-world /usr/local/bin/
ENTRYPOINT ["hello-world"]
