FROM rust:1-alpine as build

WORKDIR /app
COPY ./$ACTION .

RUN cargo build --release

FROM alpine

ARG ACTION
COPY --from=build /app/target/release/$ACTION /app

ENTRYPOINT ["/app"]
