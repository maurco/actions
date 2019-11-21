FROM rust:1.39-alpine
COPY . .
RUN cargo build --release
ENTRYPOINT ["target/release/hello-world"]
