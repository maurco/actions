FROM rust:1-alpine as build

ARG ACTION
COPY $ACTION /app
WORKDIR /app

RUN cargo build --release

FROM alpine

ARG ACTION
COPY --from=build /app/target/release/$ACTION /app

ENTRYPOINT ["/app"]
