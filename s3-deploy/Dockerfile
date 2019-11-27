FROM golang:1-alpine AS build

WORKDIR /app
COPY . .
RUN go build -v -o s3-deploy ./...

FROM alpine

COPY --from=build /app/s3-deploy /usr/local/bin
ENTRYPOINT ["s3-deploy"]
