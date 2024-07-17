#
# Arkeo
#

ARG GO_VERSION="1.20"

#
# Build
#
FROM golang:${GO_VERSION} as builder

ARG GIT_VERSION
ARG GIT_COMMIT

ENV GOBIN=/go/bin
ENV GOPATH=/go
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go install github.com/jackc/tern/v2@latest

# Download go dependencies
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire codebase and build the application
COPY . .
ARG TAG=testnet
RUN make install

#
# Main
#
FROM ubuntu:lunar

# hadolint ignore=DL3008,DL4006
RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
      jq curl htop vim ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    update-ca-certificates


# Copy the compiled binaries over.
COPY --from=builder /go/bin/sentinel /go/bin/arkeod /go/bin/indexer /go/bin/api /go/bin/tern /usr/bin/
COPY scripts /scripts

ARG TAG=testnet
ENV NET=$TAG


ENTRYPOINT ["scripts/genesis.sh"]

# default to fullnode
CMD ["arkeod", "start"]

# Health check to ensure the container is running properly
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s \
  CMD curl -f http://localhost:26657/status || exit 1