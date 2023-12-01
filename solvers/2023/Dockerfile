# 1. This tells docker to use the Rust official image
FROM rust:1.70

# install protobuf
RUN apt-get update && apt-get install -y protobuf-compiler libprotobuf-dev

WORKDIR /app

# 2. Copy the files in your machine to the Docker image
COPY . .

# Build your program for release
RUN cargo build --bin server --release

EXPOSE 8080

# Run the binary
CMD ["./target/release/server"]