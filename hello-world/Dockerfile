FROM golang:1-alpine AS build

WORKDIR /app
COPY . .
RUN go build -v -o hello-world ./...

FROM alpine

COPY --from=build /app/hello-world /usr/local/bin
ENTRYPOINT ["hello-world"]
