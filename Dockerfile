# Build application
FROM golang:1.25 AS builder

WORKDIR /production-process

COPY . .
RUN CGO_ENABLED=0 go build


# create final image

FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /
COPY --from=builder --chown=1000:1000 /production-process/compliance-as-code-attestors /app/
# CMD ["/app/./compliance-as-code-attestors "]
