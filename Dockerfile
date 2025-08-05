# syntax=docker/dockerfile:1.4

ARG GO_VERSION=1.23.11
ARG ALPINE_VERSION=3.22

########################
# --- Build Stage --- #
########################
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

# Install build dependencies
RUN apk add --no-cache git gcc g++ make

WORKDIR /src

# Pre-copy go mod files for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Optional: Set build-time variables
ARG TARGETOS
ARG TARGETARCH

# Build the Go application
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH make build

########################
# --- Final Stage --- #
########################
FROM alpine:${ALPINE_VERSION} AS final

# Add minimal runtime dependencies (curl only if your app needs it)
RUN apk add --no-cache curl

# Set working directory
WORKDIR /app

# Copy the compiled binary
COPY --from=builder /src/auto-go-app /app/auto-go-app

# Add metadata labels
LABEL org.opencontainers.image.source="https://your.repo.url" \
    org.opencontainers.image.authors="Hao Le <thienhaole92@gmail.com>" \
    org.opencontainers.image.version="${GO_VERSION}" \
    org.opencontainers.image.description="An app built with Go ${GO_VERSION}"

# Use a non-root user if security matters
RUN adduser -D appuser
USER appuser

# Use ENTRYPOINT for better shell signal handling
ENTRYPOINT ["/app/auto-go-app"]
