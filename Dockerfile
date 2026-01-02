# Build application
FROM golang:1.25 AS builder

WORKDIR /production-process

COPY . .
RUN CGO_ENABLED=0 go build -o prAttestation


# create final image

FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /
COPY --from=builder /production-process/prAttestation /usr/local/bin/prAttestation
RUN chmod +x /usr/local/bin/prAttestation

ENTRYPOINT ["/usr/local/bin/prAttestation"]
