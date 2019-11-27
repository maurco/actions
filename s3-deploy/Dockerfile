FROM golang:1-alpine as build

WORKDIR /app
COPY src src
COPY go.mod go.sum ./

RUN go build -v -o s3-deploy ./...

FROM alpine

COPY --from=build /app/s3-deploy /usr/local/bin/
ENTRYPOINT ["s3-deploy"]
