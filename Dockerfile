# Build application
FROM golang:1.25 AS builder
# Keep the build context under the module path so relative imports work.
WORKDIR /app/compliance-as-code-attestors

COPY . .
RUN CGO_ENABLED=0 go build -o compliance-as-code-attestor


# create final image

FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /
COPY --from=builder /app/compliance-as-code-attestors/compliance-as-code-attestor /usr/local/bin/compliance-as-code-attestor
RUN chmod +x /usr/local/bin/compliance-as-code-attestor

# ENTRYPOINT ["/usr/local/bin/compliance-as-code-attestor"]
