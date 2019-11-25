FROM rust:1-alpine as build

WORKDIR /app
COPY . .
RUN cargo build --release

FROM alpine

COPY --from=build /app/target/release/hello-world /usr/local/bin
ENTRYPOINT ["hello-world"]
