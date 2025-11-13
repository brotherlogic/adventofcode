# This is a multi-stage Dockerfile for an Erlang gRPC application using rebar3.

# --- Build Stage ---
# We use a specific Erlang version for reproducibility.
# You can change '25.3' to match your required Erlang/OTP version.
FROM erlang:25.3 AS builder

# Set the working directory inside the container.
WORKDIR /app

# Install rebar3, the Erlang build tool.
RUN wget https://s3.amazonaws.com/rebar3/rebar3 && chmod +x rebar3

# Copy the rebar configuration and lock file.
# This allows us to cache the dependency fetching layer.
COPY rebar.config rebar.lock ./

# Fetch the dependencies defined in rebar.config.
# The 'prod' profile is used to indicate a production build.
RUN ./rebar3 as prod get-deps

# Copy the rest of your project's source code.
# This includes your 'priv' directory and any 'src' directory you create.
COPY . .

# Compile the application and create a release.
# This will compile your .proto files and Erlang source code.
# IMPORTANT: Assumes your application is named 'aoc_server'.
# If your app has a different name, change it here.
RUN ./rebar3 as prod release -n aoc_server

# --- Final Stage ---
# We use a smaller, more secure base image for the final container.
# Alpine is a good choice for its small size.
FROM alpine:3.17

# Set the working directory for the final image.
WORKDIR /app

# Copy the compiled Erlang release from the 'builder' stage.
# The release is self-contained and includes the Erlang runtime.
# IMPORTANT: Assumes your application is named 'aoc_server'.
# If your app has a different name, change it here.
COPY --from=builder /app/_build/prod/rel/aoc_server .

# The command to start the Erlang server.
# This executes the 'foreground' command from the generated boot script.
# It runs the Erlang VM in the foreground, which is standard for containers.
CMD ["bin/aoc_server", "foreground"]
