FROM rust:1-alpine as build

WORKDIR /app
COPY . .
RUN cargo build --release

FROM alpine

COPY --from=build /app/target/release/hello-world /app
ENTRYPOINT ["/app"]
