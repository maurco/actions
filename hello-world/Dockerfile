FROM rust:1.39-alpine
WORKDIR /app
COPY . .
RUN cargo build --release
ENTRYPOINT ["/app/target/release/hello-world"]
